package go_enhance

import "fmt"

// 生产者：生成整数并发送到通道
func producer1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i // 将整数发送到通道
	}
	close(ch) // 关闭通道，表示生产者发送完成
}

// 消费者：接收并打印通道中的整数
func consumer1(ch chan int) {
	for num := range ch { // 从通道中接收整数
		fmt.Println(num)
	}
}

func Task5() {
	ch := make(chan int) // 创建一个通道

	// 启动生产者协程
	go producer1(ch)
	// 启动消费者协程
	go consumer1(ch)

	// 给协程一些时间执行
	// 在实际应用中，可以通过 sync.WaitGroup 等机制更优雅地等待
	fmt.Scanln() // 等待用户输入，确保主函数不提前退出
}
