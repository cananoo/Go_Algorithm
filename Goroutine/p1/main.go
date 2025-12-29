package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
使用go实现1000个并发控制并设置执行超时时间1秒
*/

// 假设这是一个不可中断的、耗时的实际业务逻辑
// 例如：复杂的加密运算、没有提供 Context 接口的第三方库调用
func doActualWork() string {
	// 模拟耗时：随机 0-2 秒
	cost := time.Duration(rand.Intn(2000)) * time.Millisecond
	time.Sleep(cost)
	return "业务处理结果"
}

func doTask(ctx context.Context, id int) error {
	// 1. 创建一个带缓冲的 channel 用于接收结果
	// 【重要】必须给 buffer 1，防止 Goroutine 泄漏（下面会解释）
	resultChan := make(chan string, 1)

	// 2. 开启一个新的 goroutine 去执行真正的业务逻辑
	go func() {
		// 执行耗时任务
		res := doActualWork()
		// 任务完成后，将结果发送到 channel
		resultChan <- res
	}()

	// 3. 使用 select 同时监听 结果通道 和 上下文超时
	select {
	case res := <-resultChan:
		// 路径 A: 业务在超时前完成
		fmt.Printf("任务 %d: 成功拿到结果: %s\n", id, res)
		return nil

	case <-ctx.Done():
		// 路径 B: 超时或被取消
		// 注意：这里的 return 只是退出了 doTask 函数
		// 上面启动的那个匿名 goroutine 如果还在运行，是不会被强制杀死的（除非主程退出）
		return ctx.Err()
	}
}

func main() {
	var wg sync.WaitGroup
	// 假设我们要限制并发数（可选），如果不限制直接循环即可
	// 这里为了简单，直接演示 1000 个并发启动
	total := 1000

	start := time.Now()

	for i := 0; i < total; i++ {
		wg.Add(1)

		// 【关键点】：go func 必须在循环里，让任务异步执行
		go func(id int) {
			defer wg.Done()

			// 【重点】：在这里创建 context！
			// 每次循环都会创建一个全新的 ctx，拥有独立的 1秒 倒计时
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel() // 任务结束（无论成功或超时）都要释放资源

			// 调用你的 doTask，它会阻塞直到完成或超时
			err := doTask(ctx, id)
			if err != nil {
				// 只是打印错误，不要 panic
				// fmt.Printf("任务 %d 失败: %v\n", id, err)
			}
		}(i)
	}

	wg.Wait() // 等待所有 1000 个协程结束
	fmt.Printf("所有任务处理完毕，总耗时: %v\n", time.Since(start))
}
