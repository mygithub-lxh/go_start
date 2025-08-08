package go_enhance

import (
	"fmt"
	"sync"
)

var counter1 int
var mu sync.Mutex // 创建一个 Mutex 变量用于同步

// 增加计数器的函数
func increment1(wg *sync.WaitGroup) {
	defer wg.Done() // 通知 WaitGroup 完成一个任务

	mu.Lock()   // 上锁
	counter1++  // 对计数器递增
	mu.Unlock() // 解锁
}

func Task7() {
	var wg sync.WaitGroup

	// 启动 10 个协程，每个协程执行 1000 次递增
	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个协程，WaitGroup 计数加 1
		go increment1(&wg)
	}

	// 等待所有协程执行完毕
	wg.Wait()

	// 打印最终的计数器值
	fmt.Println("Final counter value:", counter1)
}
