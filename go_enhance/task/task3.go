package go_enhance

import "fmt"

// 定义 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	width, height float64
}

// Circle 结构体
type Circle struct {
	radius float64
}

// 实现 Shape 接口的方法：Rectangle 的 Area 和 Perimeter
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// 实现 Shape 接口的方法：Circle 的 Area 和 Perimeter
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func Task3() {
	// 创建 Rectangle 实例
	r := Rectangle{width: 5, height: 3}
	fmt.Println("Rectangle:")
	fmt.Printf("Area: %.2f\n", r.Area())
	fmt.Printf("Perimeter: %.2f\n\n", r.Perimeter())

	// 创建 Circle 实例
	c := Circle{radius: 4}
	fmt.Println("Circle:")
	fmt.Printf("Area: %.2f\n", c.Area())
	fmt.Printf("Perimeter: %.2f\n", c.Perimeter())
}
