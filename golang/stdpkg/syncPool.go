package main

// sync.Pool临时对象池解决了哪些问题: http://blog.cyeam.com/golang/2017/02/08/go-optimize-slice-pool
// 前置条件: 分配slice若不指定容量则会在使用时指数级扩容（0,1,2,4,8）,故通常会指定容量
// 前置条件2: 进程获得的内存, 会被分为栈和堆，其中栈是函数调用时申请的空间, 其中的非返回值的临时变量会在执行完成后被释放(不触发GC)。
// 前置条件3: 内存碎片话, 进程获取一大块内存并分成小部分用链表链接, 每次触发申请内存操作时会从头开始遍历并将符合条件的地址返回,当申请次数频繁且大小不定的时候导致内存分布碎片化
// sync.Pool解决了这个问题
//
// 故 函数中申请的占较大内存(包括编译时不定长度的，即长度是变量)的对象以及作为返回值的对象会被申请到堆上，会在gc时被触发
//
// sync.Pool就是一个线程安全的对象缓冲池，可以获取和回收对象，缓冲池的大小会自动管理，避免过多的对象占用过多的内存。通过对缓冲池中对象的重复使用，以达到减少GC压力提高性能的目的。
// 如果应用场景中需要分配大量对象短暂使用之后就抛弃，就比较适合使用sync.Pool。
// 注意，缓冲池只用于内存对象，不能用于文件和网络连接等资源，而且获取到的对象有可能是没有初始化的。
// 临时对象池的对象适合于状态无关的对象(或者放入前或者取出后手动reset初始化)

// 关于线程安全
// 如何在多个 goroutine 之间使用同一个 pool 做到高效呢？官方的做法就是尽量减少竞争，因为 sync.pool 为每个 P（MPG模型）都分配了一个子池
// 当执行一个 pool 的 get 或者 put 操作的时候都会先把当前的 goroutine 固定到某个P的子池上面，然后再对该子池进行操作。每个子池里面有一个私有对象和共享列表对象，私有对象是只有对应的 P 能够访问，因为一个 P 同一时间只能执行一个 goroutine，因此对私有对象存取操作是不需要加锁的。共享列表是和其他 P 分享的，因此操作共享列表是需要加锁的。
//
// 获取对象过程是：
// 1）固定到某个 P，尝试从私有对象获取，如果私有对象非空则返回该对象，并把私有对象置空；
// 2）如果私有对象是空的时候，就去当前子池的共享列表获取（需要加锁）；
// 3）如果当前子池的共享列表也是空的，那么就尝试去其他P的子池的共享列表偷取一个（需要加锁）；
// 4）如果其他子池都是空的，最后就用用户指定的 New 函数产生一个新的对象返回。
//
// 可以看到一次 get 操作最少 0 次加锁，最大 N（N 等于 MAXPROCS）次加锁。
//
// 归还对象的过程：
// 1）固定到某个 P，如果私有对象为空则放到私有对象；
// 2）否则加入到该 P 子池的共享列表中（需要加锁）。
// 可以看到一次 put 操作最少 0 次加锁，最多 1 次加锁。
//
// 由于 goroutine 具体会分配到那个 P 执行是 golang 的协程调度系统决定的，因此在 MAXPROCS>1 的情况下，多 goroutine 用同一个 sync.Pool 的话，各个 P 的子池之间缓存的对象是否平衡以及开销如何是没办法准确衡量的。但如果 goroutine 数目和缓存的对象数目远远大于 MAXPROCS 的话，概率上说应该是相对平衡的。

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// 一个[]byte的对象池，每个对象为一个[]byte
// Get方法和Put方法
var bytePool = sync.Pool{
	// 当池里没有数据的时候会根据方法new一个
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func TestNewWithoutPool(b *testing.B) {
	// 不使用对象池
	for i := 0; i < b.N; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
}

func TestNewWithPool(b *testing.B) {
	// 使用对象池// 会快些
	for i := 0; i < b.N; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
}

func main() {
	a := bytePool.Get().(*[]byte)
	(*a)[0] = 1
	b := bytePool.Get().(*[]byte)
	(*b)[0] = 1
	(*b)[1] = 1
	bytePool.Put(b)
	bytePool.Put(a)
	// runtime.GC() // 主动触发GC
	c := bytePool.Get().(*[]byte)
	fmt.Println("从放入a, b的池中取出的c", (*c)[:10]) // 还有原来的值, 故临时对象池的对象适合于状态无关的对象
	fmt.Println("New without pool:")
	fmt.Println(testing.Benchmark(TestNewWithoutPool))
	fmt.Println("New with pool:")
	fmt.Println(testing.Benchmark(TestNewWithPool))
}
