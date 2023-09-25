package main
import "fmt"
func main() {
	var  age int =10
	switch age  {
	case 10:
		fmt.Println("10")
		fallthrough  //默认只能穿透一层
	case 20:
		fmt.Println("20")
		fallthrough  //默认只能穿透一层
	case 30:
		fmt.Println("30")
	case 40:
		fmt.Println("40")

	}
}




