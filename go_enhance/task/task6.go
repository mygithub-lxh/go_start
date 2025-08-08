package go_enhance

import "fmt"

// 生产者：向带缓冲的通道发送整数
func producer2(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i // 发送整数到通道
	}
	close(ch) // 关闭通道，表示生产者发送完成
}

// 消费者：接收并打印通道中的整数
func consumer2(ch chan int) {
	for num := range ch { // 从通道中接收整数
		fmt.Println(num)
	}
}

func Task6() {
	ch := make(chan int, 10) // 创建一个带缓冲区大小为 10 的通道

	// 启动生产者协程
	go producer2(ch)
	// 启动消费者协程
	go consumer2(ch)

	// 给协程一些时间执行
	fmt.Scanln() // 等待用户输入，确保主函数不提前退出
}
