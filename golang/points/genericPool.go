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

type closerWithDeadLine struct {
	closer io.Closer
	dl     time.Time
}

type GenericPool struct {
	sync.Mutex
	pool        chan closerWithDeadLine
	maxOpen     int  // 池中最大资源数
	numOpen     int  // 当前池中资源数
	minOpen     int  // 池中最少资源数
	closed      bool // 池是否已关闭
	maxLifetime time.Duration
	factory     factory // 创建连接的方法
}

func NewGenericPool(minOpen, maxOpen int, factory factory, maxLifetime time.Duration) (*GenericPool, error) {
	if maxOpen <= 0 || minOpen > maxOpen {
		return nil, ErrInvalidConfig
	}
	// 初始化池结构
	p := &GenericPool{
		maxOpen:     maxOpen,
		minOpen:     minOpen,
		factory:     factory,
		maxLifetime: maxLifetime,
		pool:        make(chan closerWithDeadLine, maxOpen),
	}
	deadLine := time.Now().Add(p.maxLifetime)
	// 建立最小连接数
	for i := 0; i < minOpen; i++ {
		closer, err := factory()
		if err != nil {
			continue
		}
		p.numOpen++
		p.pool <- closerWithDeadLine{closer: closer, dl: deadLine}
	}
	return p, nil
}

func (p *GenericPool) Acquire() (io.Closer, error) {
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
	for {
		select {
		case c := <-p.pool:
			if p.numOpen > p.minOpen && time.Now().After(c.dl) {
				p.Close(c.closer)
				continue
			}
			return c.closer, nil
		default:
			p.Lock()
			// 超出连接上限后, 等待池中有连接被归还
			if p.numOpen >= p.maxOpen {
				c := <-p.pool
				p.Unlock()
				return c.closer, nil
			}
			// 新建连接
			closer, err := p.factory()
			if err != nil {
				p.Unlock()
				return nil, err
			}
			p.numOpen++
			p.Unlock()
			return closer, nil
		}
	}
}

// 释放单个资源到连接池
func (p *GenericPool) Release(closer io.Closer) error {
	if p.closed {
		return ErrPoolClosed
	}
	p.pool <- closerWithDeadLine{closer: closer, dl: time.Now().Add(p.maxLifetime)}
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
	for c := range p.pool {
		c.closer.Close()
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
	pool, _ := NewGenericPool(3, 8, OpenFile, 20*time.Second)
	wg := sync.WaitGroup{}
	length := 10
	for i := 0; i < length; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func(i int) {
			file, _ := pool.Acquire()
			fmt.Println("Open file ", i+1, " times: ", file, " numsOpen in pool: ", pool.numOpen)
			time.Sleep(10 * time.Second)
			pool.Release(file)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
