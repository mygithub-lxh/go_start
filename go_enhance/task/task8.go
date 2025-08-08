package go_enhance

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter2 int64 // 使用 int64 类型作为计数器，原子操作需要是整型数据

// 增加计数器的函数
func increment2(wg *sync.WaitGroup) {
	defer wg.Done() // 通知 WaitGroup 完成一个任务

	atomic.AddInt64(&counter2, 1) // 使用原子操作递增计数器
}

func Task8() {
	var wg sync.WaitGroup

	// 启动 10 个协程，每个协程执行 1000 次递增
	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个协程，WaitGroup 计数加 1
		go increment2(&wg)
	}

	// 等待所有协程执行完毕
	wg.Wait()

	// 打印最终的计数器值
	fmt.Println("Final counter value:", counter2)
}
