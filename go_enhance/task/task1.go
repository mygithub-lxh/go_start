package go_enhance

import (
	"fmt"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func addTen(a *int) {
	*a += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func doubleSlice(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

// 俩个线程分别输出奇数和偶数
func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
		time.Sleep(100 * time.Millisecond) // 模拟任务延迟
	}
}

func printEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数:", i)
		time.Sleep(100 * time.Millisecond) // 模拟任务延迟
	}
}

func Task1() {
	value := 5
	fmt.Println("原值：", value)
	addTen(&value) // 传入 value 的地址
	fmt.Println("修改后：", value)

	data := []int{1, 2, 3, 4}
	fmt.Println("原切片：", data)

	doubleSlice(&data) // 传入切片的指针

	fmt.Println("修改后：", data)
	go printOdd()
	go printEven()
	time.Sleep(2 * time.Second)

}
