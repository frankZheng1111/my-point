// https://studygolang.com/articles/12333
package main

import (
	"errors"
	"io"
	"sync"
	"time"
)

import (
	"fmt"
	"os"
	"runtime"
)

var (
	ErrInvalidConfig = errors.New("invalid pool config")
	ErrPoolClosed    = errors.New("pool closed")
)

type factory func() (io.Closer, error)

type Pool interface {
	Acquire() (io.Closer, error) // 获取资源
	Release(io.Closer) error     // 释放资源
	Close(io.Closer) error       // 关闭资源
	Shutdown() error             // 关闭池
}

type GenericPool struct {
	sync.Mutex
	pool    chan io.Closer
	hbChan  chan struct{}
	maxOpen int     // 池中最大资源数
	numOpen int     // 当前池中资源数
	minOpen int     // 池中最少资源数
	closed  bool    // 池是否已关闭
	factory factory // 创建连接的方法
}

func (p *GenericPool) HeartBeatCloserHandle() {
	for {
		select {
		case _ = <-p.hbChan:
			continue
		case _ = <-time.After(1 * time.Second):
			if p.numOpen > p.minOpen {
				select {
				case closer := <-p.pool:
					p.Close(closer)
					continue
				default:
					continue
				}
				continue
			}
		}
	}
}

func (p *GenericPool) HeartBeat() {
	p.hbChan <- struct{}{}
}

func NewGenericPool(minOpen, maxOpen int, factory factory) (*GenericPool, error) {
	if maxOpen <= 0 || minOpen > maxOpen {
		return nil, ErrInvalidConfig
	}
	// 初始化池结构
	p := &GenericPool{
		maxOpen: maxOpen,
		minOpen: minOpen,
		factory: factory,
		pool:    make(chan io.Closer, maxOpen),
		hbChan:  make(chan struct{}),
	}
	// 建立最小连接数
	for i := 0; i < minOpen; i++ {
		closer, err := factory()
		if err != nil {
			continue
		}
		p.numOpen++
		p.pool <- closer
	}
	go p.HeartBeatCloserHandle()
	return p, nil
}

func (p *GenericPool) Acquire() (io.Closer, error) {
	p.HeartBeat()
	if p.closed {
		return nil, ErrPoolClosed
	}
	for {
		closer, err := p.getOrCreate()
		if err != nil {
			return nil, err
		}
		return closer, nil
	}
}

func (p *GenericPool) getOrCreate() (io.Closer, error) {
	select {
	case closer := <-p.pool:
		return closer, nil
	default:
		p.Lock()
		defer p.Unlock()
		// 超出连接上限后, 等待池中有连接被归还
		if p.numOpen >= p.maxOpen {
			closer := <-p.pool
			return closer, nil
		}
		// 新建连接
		closer, err := p.factory()
		if err != nil {
			return nil, err
		}
		p.numOpen++
		return closer, nil
	}
}

// 释放单个资源到连接池
func (p *GenericPool) Release(closer io.Closer) error {
	p.HeartBeat()
	if p.closed {
		return ErrPoolClosed
	}
	p.pool <- closer
	return nil
}

// 关闭单个资源
func (p *GenericPool) Close(closer io.Closer) error {
	p.Lock()
	closer.Close()
	p.numOpen--
	p.Unlock()
	return nil
}

// 关闭连接池，释放所有资源
func (p *GenericPool) Shutdown() error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	close(p.pool)
	for closer := range p.pool {
		closer.Close()
		p.numOpen--
	}
	p.closed = true
	p.Unlock()
	return nil
}

// 测试运行

func OpenFile() (io.Closer, error) {
	file, err := os.Open("./README.md")
	return file, err
}

func main() {
	runtime.GOMAXPROCS(1)
	pool, _ := NewGenericPool(3, 8, OpenFile)
	wg := sync.WaitGroup{}
	length := 10
	for i := 0; i < length; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func(i int) {
			file, _ := pool.Acquire()
			fmt.Println("Open file ", i+1, " times: ", file, " numsOpen in pool: ", pool.numOpen)
			time.Sleep(time.Duration(length-i) * time.Second)
			pool.Release(file)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("current numOpen %d\r", pool.numOpen)
	}
}
