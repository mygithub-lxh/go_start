package go_enhance

import "fmt"

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Employee 结构体，组合 Person 结构体
type Employee struct {
	Person     // 嵌入 Person 结构体
	EmployeeID string
}

// 为 Employee 实现 PrintInfo 方法
func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s\nAge: %d\nEmployeeID: %s\n", e.Name, e.Age, e.EmployeeID)
}

func Task4() {
	// 创建 Employee 实例
	emp := Employee{
		Person:     Person{Name: "John", Age: 30},
		EmployeeID: "E12345",
	}

	// 调用 PrintInfo 方法
	emp.PrintInfo()
}
