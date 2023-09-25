package main
import "fmt"
func main() {
	var age int 
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	if age>18 && age<=60 { // 这个大括号不能换行的，换行就会报错的
		fmt.Println("成年人")
	}else if age>60{
		fmt.Println("老年人");
	}else{
		fmt.Println("未成年人");
	}
}




