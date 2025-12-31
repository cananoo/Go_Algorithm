package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
sync.Cond实现多生产者多消费者
核心设计思路
为了最高效地实现，我们需要 一把锁 (sync.Mutex) 和 两个条件变量 (sync.Cond)：
	1. lock: 保护共享队列（Slice）。
	2. notFull (Cond): 生产者在这里等待。如果队列满了，生产者就睡在这个条件上。
       notEmpty (Cond): 消费者在这里等待。如果队列空了，消费者就睡在这个条件上。
*/

// FIFO 是一个线程安全的队列
type FIFO struct {
	lock     sync.Mutex
	notFull  *sync.Cond // 条件：队列不满（生产者关注）
	notEmpty *sync.Cond // 条件：队列不空（消费者关注）

	data     []int // 共享数据容器
	capacity int   // 队列最大容量
}

// NewFIFO 初始化队列
func NewFIFO(capacity int) *FIFO {
	q := &FIFO{
		data:     make([]int, 0, capacity),
		capacity: capacity,
	}
	// ⚠️ 关键点：两个 Cond 共用同一把锁 q.lock
	q.notFull = sync.NewCond(&q.lock)
	q.notEmpty = sync.NewCond(&q.lock)
	return q
}

// Produce 生产者方法
func (q *FIFO) Produce(item int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// 1. 【循环检查】如果队列已满，生产者必须等待
	// ⚠️ 必须用 for 而不是 if，防止“虚假唤醒”
	for len(q.data) >= q.capacity {
		fmt.Println("队列满，生产者等待...")
		q.notFull.Wait() // 阻塞，并自动释放锁；被唤醒时重新获取锁
	}

	// 2. 加入数据
	q.data = append(q.data, item)
	fmt.Printf("生产: %d | 当前队列: %v\n", item, q.data)

	// 3. 唤醒一个消费者 (队列现在肯定不空了)
	// 使用 Signal 即可，不需要 Broadcast，因为只需唤醒一个干活的
	q.notEmpty.Signal()
}

// Consume 消费者方法
func (q *FIFO) Consume() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	// 1. 【循环检查】如果队列为空，消费者必须等待
	for len(q.data) == 0 {
		fmt.Println("队列空，消费者等待...")
		q.notEmpty.Wait()
	}

	// 2. 取出数据 (模拟队列头部弹出)
	item := q.data[0]
	q.data = q.data[1:]
	fmt.Printf("消费: %d | 当前队列: %v\n", item, q.data)

	// 3. 唤醒一个生产者 (队列现在肯定不满了，有空位了)
	q.notFull.Signal()

	return item
}

func main() {
	// 创建一个容量为 3 的队列
	queue := NewFIFO(3)
	var wg sync.WaitGroup

	// 启动 5 个生产者
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 2; j++ { // 每个生产者生产 2 个
				num := id*10 + j
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				queue.Produce(num)
			}
		}(i)
	}

	// 启动 2 个消费者
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ { // 每个消费者消费 5 个
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				queue.Consume()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("程序结束")
}
