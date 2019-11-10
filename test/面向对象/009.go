package main

import "fmt"

type Rectangle struct {
	width, height float64
	color         string
}

// 如果声明接收者为指定，当使用T类型来调用时，go会自动转换为*T，太TM聪明了
func (r *Rectangle) SetColor(color string) {
	r.color = color
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 2, "白色"}
	r2 := Rectangle{9, 4, "蓝色"}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())

	fmt.Println("Color of r2 is: ", r2.color)
	fmt.Println("set new color")
	r2.SetColor("go色") // 等价于(&r2).SetColor("绿色"),这里r2不需要传地址，当然，传地址也不错，只是go会自动帮助转换，没有必要
	fmt.Println("Color of r2 is: ", r2.color)
	fmt.Println(r2)
}
