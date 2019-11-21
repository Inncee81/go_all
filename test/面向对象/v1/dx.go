package main
import (
	"fmt"
)
type Real struct {
	width, height float64
	color         string
}
type GReal struct{
	real Real
	height float64
}

func (r *Real) SetColor(color string) {
	r.color = color
}
func (r Real) GetColor()(color string) {
	return r.color
}

func (r Real) area() float64 {
	return r.height * r.width
}
func (r GReal) garea() float64 {
return	r.real.height*r.real.width*r.height
}

func main() {
	r1 := Real{12, 2, "白色"}
	r2 := Real{19, 12, "黑色"}
	r3:=GReal{r1,10}
	r1.SetColor("白白色")
	r2.SetColor("黑黑色")
	fmt.Println(	r1.area())
	fmt.Println( r2.area())
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r1.GetColor())
	fmt.Println(r3.garea())


}
