package main
import "fmt"

func main() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("     %d x %d = %d",j,i, i*j);
		}
		fmt.Println()
	}
}




// 打印金字塔
// 	   *        1-> 1个*   i+(i-1) -> 2i-1    空格 2   3-1   空格数=n-层
//    ***		2-> 3个*	 2*2-1=3			 1	 3-2
//   *****		3-> 5个*	 2*3-1=5			 0   3-3
// func main() {
// 	var count int =30
// 	for i := 1; i <= count; i++ {
// 		// 打印空格
// 		for k := 0; k < count-i; k++ {
// 			fmt.Print(" ");
// 		}
// 		//打印星星
// 		xx:=2*i-1
// 		for j := 1; j <= xx; j++ {
// 			//fmt.Print("*"); // 实心金字塔
// 			if j==1 || j==xx { // 空心金字塔
// 				fmt.Print("*");  //第一个和最后一个打印星星
// 			}else{
// 				if i==count{ // 如果是最后一行
// 					fmt.Print("*");
// 				}else{
// 					fmt.Print(" ");
// 				}
// 			}
// 		}
// 		fmt.Println(); // 换行
// 	}
// }

// for range
// func main() {
// 	var str string="hello world!"
// 	for i := 0; i < len(str); i++ {
// 		fmt.Printf("%c \n",str[i]);
// 	}

// 	str="abc 西安"
// 	for index, val := range str {
// 		fmt.Printf("index=%d, val=%c \n",index,val)
// 	}

// }



// func main() {
// 	// 配合break 用法
// 	for {
// 		fmt.Println("您好！",i)
// 	}
// }



// func main() {
// 	i:=0
// 	for  i < 10 {
// 		fmt.Println("您好！",i)
// 		i++
// 	}
// }



// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("您好！",i)
// 	}
// }







