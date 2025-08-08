package go_enhance

import (
	"fmt"
	"time"
)

type Task func()

func runTaskWithTiming(task Task, taskName string) {
	startTime := time.Now()

	// 执行任务
	task()

	// 统计任务执行时间
	duration := time.Since(startTime)
	fmt.Printf("任务 %s 执行时间: %v\n", taskName, duration)
}
func downloadFile() {
	fmt.Println("开始下载文件...")
	time.Sleep(2 * time.Second) // 模拟下载时间
	fmt.Println("文件下载完成!")
}

// 任务示例2：模拟处理数据
func processData() {
	fmt.Println("开始处理数据...")
	time.Sleep(1 * time.Second) // 模拟数据处理时间
	fmt.Println("数据处理完成!")
}
func Task2() {

	// 定义任务
	tasks := []struct {
		task     Task
		taskName string
	}{
		{downloadFile, "下载文件"},
		{processData, "处理数据"},
	}

	// 启动多个协程并行执行任务
	for _, t := range tasks {
		go runTaskWithTiming(t.task, t.taskName)
	}

	// 等待足够的时间让所有任务执行完毕
	time.Sleep(3 * time.Second)

}
