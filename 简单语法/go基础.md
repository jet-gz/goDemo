[TOC]



#  坏境

## 安装SDK
[下载地址](https://golang.google.cn/dl/) 
### windows 安装
可以直接下载安装msi  不需要配置坏境变量默认配置好了  
zip 里面包含源码，但是需要配置坏境变量

### liunx安装SDK
```shell

   wget https://dl.google.com/go/go1.20.2.linux-amd64.tar.gz
   tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
   cd /usr/local
   cd go # 发现里面右go目录
   ## 可以看到结果
   bin/go version  #
```
 配置 GOROOT 和 PATH坏境变量  
    GOROOT： export GOROOT=/usr/local/go  
    PATH ：  export PATH=$PATH:$GOROOT/bin:$GOBIN
vim /root/.profile   最后加入到  etc/.profile 
将2个环境变量加入到最下面  
```go
export GOROOT=/usr/local/go    
export PATH=$PATH:$GOROOT/bin:$GOBIN  
```
刷新坏境变量  
source /etc/profile    
root   source /root/.profile

## 第一次使用问题
go mod init 之类的错误 可以执行下命令  
go env -w GO111MODULE=off

## 坏境变量配置
1. GOROOT 指定sdk 路径    新建系统坏境变量  GOROOT = go的安装目录C:\Program Files\Go  
2. Path 添加sdk bin目录   修改PATH    ;%GOROOT%/bin
3. GOPATH 工作目录

Path 环境变量添加只需要 引用 sdk的变量  %sdk变量%/bin


# 常量
## 简单常量
```go
const Pi=3.1415926
const abc="abc"
```
## 声明赋值常量
```go
const beef, two, sundsewew = "eat", 2, "veg"
const (monday,tuesday,wenday,t=1,2,3,4);
```
## 用作枚举常量
```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```
### 枚举简写
    iota用法  
```go
        package main
        import "fmt"
        func main() {

            const (
                a = iota   //0
                b          //1
                c          //2
                d = "ha"   //独立值，iota += 1
                e          //"ha"   iota += 1
                f = 100    //iota +=1
                g          //100  iota +=1
                h = iota   //7,恢复计数
                i          //8
            )
            fmt.Println(a,b,c,d,e,f,g,h,i)
        }
```

    iota，特殊常量，可以认为是一个可以被编译器修改的常量。  
    iota 是0 ,后面都会依次，0,1,2,简单来说每次遇见const iota 都会重置0  
```go
const (
    a=iota
    b
    c
)
```
### 某个类型作为枚举常量的类型

```go
type Color int
	const (
		RED Color = iota // 0
		ORANGE // 1
		YELLOW // 2
		GREEN // ..
		BLUE
		INDIGO
		VIOLET // 6
	)
```

# 变量
    变量声明了一般必须使用
    小写开头的变量外部包是没法访问的（私有的），只有首字母大写才可以访问（公开的）
```go
var a int
var b bool
var c string
var (a1 int,b1 bool,c1 string)
```
    当变量声明后就会有默认值 int 0，float 0.0  bool false   string 为空字符串，  指针为 nil， 所有的内存在go中都需要经过初始化, 如果一个变量在函数体外定义，则它就是个全局变量，函数体内声明局部变量  简写 a:=1    不支持函数外部定义
```go
var goos string = runtime.GOOS
    fmt.Printf("The operating system is: %s\n", goos)
	// path  局部变量
    path:= os.Getenv("PATH")  
    fmt.Printf("Path is %s\n", path)
```
//  局部变量和全局变量 名称相同，  编译器会采用就近原则


# 值类型
* 内存地址都是以16进制表示
* int、float、bool、string 都是值类型，使用这些类型的变量直接指向内存中的值，像数组复合类型也属于值类型
* i=j 是将i的值拷贝了一分给了j
* 通过&i 可以获取到内存地址，值类型变量的值都存在栈中
* 引用类型，变量存的是一个地址，
* 内存的地址被称为指针，
* 指针指向的内存都是连续的，这也是计算效率最好的一种存储形式。每个字节指向了下一个字节的地址
* 引用类型赋值，只是地址拷贝


# 字符串
* 不可变：一旦赋值就不能再修改了
* 表示方式
    - 双引号 :一般常规的字符串，特殊字符串需要转移
    - 反引号 : `` 直接原分不动的输出
* 拼接方式：
    - 直接用+ 拼接，或者+=  适用于短的字符串拼接
    - 多行拼接 可以换行，换行后，需要将+放在上一行

## 解释字符串
* \n：换行符
* \r：回车符
* \t：tab 键
* \u 或 \U：Unicode 字符
* \\：反斜杠自身
## 非解释字符串
    字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。


# 打印
## 通用
```go
    %v	值的默认格式表示
    %+v	类似%v，但输出结构体时会添加字段名
    %#v	值的Go语法表示
    %T	值的类型的Go语法表示
    %%	百分号
```
## 整数
```go
    %b	表示为二进制
    %c	该值对应的unicode码值
    %d	表示为十进制
    %o	表示为八进制
    %q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
    %x	表示为十六进制，使用a-f
    %X	表示为十六进制，使用A-F
    %U	表示为Unicode格式：U+1234，等价于"U+%04X"
```
## 布尔
```go
    %t	单词true或false
```
## 浮点数与复数的两个组分
```go
    %b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
    %e	科学计数法，如-1234.456e+78
    %E	科学计数法，如-1234.456E+78
    %f	有小数部分但无指数部分，如123.456
    %F	等价于%f
    %g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
    %G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
```
## 字符串和[]byte
```go
    %s	直接输出字符串或者[]byte
    %q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
    %x	每个字节用两字符十六进制数表示（使用a-f）
    %X	每个字节用两字符十六进制数表示（使用A-F）    
```
## 指针：
```go
    %p	表示为十六进制，并加上前导的0x    
```


# 类型与运算
## 整数与浮点数
 go没有float 和doule 类型，只有 float32 和 float64
 * 整数：
    - int8（-128 -> 127）
    - int16（-32768 -> 32767）
    - int32（-2,147,483,648 -> 2,147,483,647）
    - int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
* 无符号整型
    - uint8（0 -> 255）
    - uint16（0 -> 65,535）
    - uint32（0 -> 4,294,967,295）
    - uint64（0 -> 18,446,744,073,709,551,615）

* 浮点型（IEEE-754 标准）：
    - float32（+- 1e-45 -> +- 3.4 * 1e38）
    - float64（+- 5 1e-324 -> 107 1e308）  

整数是计算最快的一种类型


# 类型转换
## Sprint
``` go
    package main
    import "fmt"

    func main() {

        var num1 int =99
        var num2 float64 =23.5566
        var b bool =true
        var str string
        
        //Sprintf根据format参数生成格式化的字符串并返回该字符串。
        str=fmt.Sprintf("%d",num1)
        //%T 打印类型
        fmt.Printf("类型 %T  str=%v \n",str,str)
        // %f 有小数部分但无指数部分，如123.456
        str=fmt.Sprintf("%f",num2)
        fmt.Printf("类型 %T  str=%v \n",str,str)
        //%t	单词true或false
        str=fmt.Sprintf("%t",b)
        fmt.Printf("类型 %T  str=%v \n",str,str)
    }

```
## strconv
```go
package main
import "fmt"
import "strconv"

func main() {

	var num1 int =99
	var num2 float64 =23.5566
	var b bool =true
	var str string
	
	str=strconv.FormatInt(int64(num1),10)
	fmt.Printf("类型 %T  str=%q \n",str,str)
	// f 格式，10 保留位数，64 是float64
	str=strconv.FormatFloat(num2,'f',10,64)
	fmt.Printf("类型 %T  str=%q \n",str,str)

	str=strconv.FormatBool(b)
	fmt.Printf("类型 %T  str=%v \n",str,str)

}
```
##  Parse
他一般有  ParseBool   ParseInt  ParseFloat 等  
返回2个值 接受方式用  n1,_=

# 指针
* 获取变量地址用&
  ``` go
    var i int=10
    fmt.Println("i的地址是",&i)
  ``
  
```go
package main
import "fmt"
// import "strconv"

func main() {
	var i int=10
	fmt.Println("i的地址是",&i)

	//p是一个指针变量， 类型为*int p本身值是 &i
	// p 存储了一个地址  0xc00000e0a8  就是i的地址，这个地址指向了i的空间
	var p *int =&i
	fmt.Printf("p=%v\n",p)
	fmt.Printf("p的地址=%v",&p)
	// 通过p 取出 i的值 10
	fmt.Printf("p指向的值是=%v \n",*p)
	// 修改了i的值
	*p=20

	fmt.Printf("p指向的值是=%v \n",*p)
	fmt.Println(i)

}

```



# 获取用户的输入
## fmt.Scanln()
```go
package main
import "fmt"
// import "strconv"

func main() {
	var name string
	var age byte
	var sal float32
	
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请输入年纪")
	fmt.Scanln(&age);

	fmt.Println("请输入薪水")
	fmt.Scanln(&sal);

	fmt.Printf("姓名%v \n 年纪%v \n  薪水%v \n ",name,age,sal );

}

```
## fmt.Scanf()
```go
package main
import "fmt"
// import "strconv"

func main() {
	var name string
	var age byte
	var sal float32

	fmt.Println("请依次输入，姓名，年纪 薪水,用空格隔开")
	//接受数据的类型可以参考文档
	fmt.Scanf("%s %d %f",&name,&age,&sal)
	fmt.Printf("姓名：%v 年龄：%v 薪资：%v ",name,age,sal)
}

```
# 运算
## 除法运算
``` go
package main
import "fmt"

func main() {
	var n1 float32 =10/4
	//这里 因为都是整数相除的，那么除后会去掉小数部分，保留整数部分,
	fmt.Println(n1)  // 结果是2
	//如果要保留小数需要  10.0/4
	var n2 float32 =10.0/4
	fmt.Println(n2)  // 结果是2
}
```
## 求模
公式 a%b =a-a/b*b
```go
	fmt.Println("10%3=",10%3,"     10-10/3*3=",10-10/3*3) //1
	fmt.Println("-10%3=",-10%3,"     -10-(-10)/3*3=",-10-(-10)/3*3) //-1
	fmt.Println("10%-3=",10%-3,"     10-10/-3*-3=",10-10/-3*-3)   //1
```
# 流程控制
# if else 语句
> go 语言中，大括号，有时候是不能换行的，换行就会报错
> go 中的if 一般是不需要加括号的，除非是，if(a>0 && a<10) || a==90 其实这里的括号也不是给if的加的，是给条件加的 
```go
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
```
## swich
> swich 语句后面是不需要加 break的
> case :后面可是常量，也可以是个函数，也可以运算

``` go
package main
import "fmt"
func main() {
	var key byte
	fmt.Println("请输入字符串a,b,c,d,e");
	fmt.Scanf("%c",&key);

	switch key {
	case 'a':
		fmt.Println("======>a");
	case 'b':
		fmt.Println("======>b");
	case 'c':
		fmt.Println("======>c");
	case 'd':
		fmt.Println("======>d");
	case 'e':
		fmt.Println("======>e");
	case 'f','g','h':
		fmt.Println("======>其他输入");
	default:
		fmt.Println("输入有误");
	}
}


package main
import "fmt"
func main() {
	var  age int =10
	switch  {  //也可以 switch age:=10;{}
	case age==10:
		fmt.Println("10")
	case age>10:
		fmt.Println("大于10")
	}
}


```
## switch 穿透 fallthrought
``` go
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

```
# 循环
## for
>例子
``` go 
package main
import "fmt"
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("您好！",i)
	}
}

```
> 格式
循环初始，循环条件，循环迭代
> 顺序
1. 初始化i=0
2. 条件 i<10
3. 如果为真，执行fmt.Println
4. 执行迭代，i++
5. 反复执行，直到条件为假 退出

## for 第二种写法
> 和第一种写法其实一样
```go
package main
import "fmt"
func main() {
	i:=0
	for  i < 10 {
		fmt.Println("您好！",i)
		i++
	}
}
```
## for 第三种写法
```go
package main
import "fmt"
func main() {
	// 配合break 用法
	for {
		fmt.Println("您好！",i)
	}
}
```

## 字符串遍历
``` go
func main() {
	var str string="hello world!"
    //str2=[]rune(str)  // 将下面的str 改成str2  这个就是切片
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n",str[i]);
	}
}

```
> 如果 str字符串包含中文，就会出问题，乱码，
> 传统对字符串的便利是按照字字节来遍历的，而一个汉字再utf8 中包含3个字节
> 解决的方法，需要将，str转换成切片， 或者用 for-range

## for-range

``` go
func main() {
	var str string="hello world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n",str[i]);
	}

	str="abcd西安"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n",index,val)
	}
}

```
> 结果  
> 
![本地图](img/20230405155727.png "range")

## 循环案例
### 打印金字塔
``` go

// 打印金字塔
// 	   *        1-> 1个*    i+(i-1) -> 2i-1    空格 2   3-1   空格数=n-层
//    ***		2-> 3个*	 2*2-1=3			  1	 3-2
//   *****		3-> 5个*	 2*3-1=5			  0   3-3
func main() {
	var count int =30
	for i := 1; i <= count; i++ {
		// 打印空格
		for k := 0; k < count-i; k++ {
			fmt.Print(" ");
		}
		//打印星星
		xx:=2*i-1
		for j := 1; j <= xx; j++ {
			//fmt.Print("*"); // 实心金字塔
			if j==1 || j==xx { // 空心金字塔
				fmt.Print("*");  //第一个和最后一个打印星星
			}else{
				if i==count{ // 如果是最后一行
					fmt.Print("*");
				}else{
					fmt.Print(" ");
				}
			}
		}
		fmt.Println(); // 换行
	}
}


```
### 九九乘法表
``` go
func main() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("     %d x %d = %d",j,i, i*j);
		}
		fmt.Println()
	}
}
```
# 随机数
``` go
package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var count int =0
	for  {
		//生成100次 退出
		if count==100 {
			break
		}
		//t:=time.Now().Unix()  // 返回当前的时间戳
		//fmt.Println(t)
		rand.Seed(time.Now().UnixNano())// 纳秒 以当前时间作为随机种子  重复的几率还是很大
		n:=rand.Intn(100)+1  // 生产1到100 的随机数
		fmt.Println(n)
		count++
	}

	
}


```
## break  continue
和C#里面的差不多
>break 直接跳出循环，结束了  
>continue 跳出当次循环，然后进行下一次，
## goto
```go
package main
import (
	"fmt"

)

func main() {

	var n int =4
	fmt.Println("1")
	fmt.Println("2")
	if(n>3){
		goto Lable
	}
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	Lable:
	fmt.Println("6")
	fmt.Println("7")
}

```
## return 
>结束方法 函数。

# 函数
## 实例
``` go
package main
import (
	"fmt"
)

func main() {

	result:=cal(2.3,2.7,'+')
	fmt.Println("result=",result)
}

func cal(n1 float64,n2 float64,o byte ) float64{

	switch o {
		case '+':
			return n1+n2
		case '-':
			return n1-n2
		case '*':
			return n1*n2
		case '/':
			return n1/n2
		default:
			fmt.Println("不支持的计算")
		return -99999
	}

} 
```
## 返回多个参数的函数
```go
package main
import (
	"fmt"
	
)

func main() {
	fmt.Println("返回多个参数的");
	
	// res1,res2:=A(2,3)
	// fmt.Println("sum=",res1)
	// fmt.Println("ss=",res2)

	_, res2:=A(2,3)
	fmt.Println("ss=",res2)
}

func A(a int,b int)(int,int){
	sum:=a+b
	ss:=a*b
	return sum ,ss
}


```

## 函数的引用传值方式
>上面的函数传值 都是以值传值的方式，通过拷贝的方式
``` go
package main
import (
	"fmt"
)

func main() {
	a:=4
	A(&a) // 直接传了a的地址
	fmt.Println("main a=",a);
}
// a的指针
func A(a *int ) {
	*a++   // 根据地址找到修改了
	fmt.Println("A() a=",*a);
}
```
## 将函数当作参数传递
```go
package main
import (
	"fmt"
)

func main() {
	
  s:=SS(sum,20,40)
   fmt.Println("s=",s);

}

func SS(dd func(int,int)int, a int,b int) int{
	return dd(a,b)
}

func sum(a int,b int ) int {
	return a+b
}

```
## 自定义类型
我们可以给类型取别名  
>demo1
``` go
package main
import (
	"fmt"
)

func main() {
	
	type myInt int
	var n1 myInt
	var n2 int
    // n1 和n2  是不同的类型

	n1=100
	//n2=n1  // 报错的， 不同的类型，除非强制转换
	n2=int(n1)
	fmt.Println("n1=",n1,"n2=",n2);

}

```
>demo2

``` go
package main
import (
	"fmt"
)

type my func(int,int)int

func main() {
  c:=SS(sum,100,200)
  fmt.Println(c);
}

func SS(dd my, a int,b int) int{
	return dd(a,b)
}

func sum(a int,b int ) int {
	return a+b
}

```
>demo3
``` go
package main
import (
	"fmt"
)

func main() {
  s,ss:=count(2,3)
  fmt.Println("sum=",s,"ss=",ss);
}

func count(a int,b int)(sum int,ss int){
	sum=a+b
	ss=a*b
	return
}


```
## 可变参数
``` go
package main
import (
	"fmt"
)

func main() {
	n:=sum(3,4,10)

	fmt.Println("sum=",n);
}

func sum(n1 int,args... int)int{

	sum:=n1
	for i := 0; i < len(args); i++ {
		sum+=args[i]
	}
	return sum
}
```

## init 函数
>介绍 
1. 每个源文件中都可以包含一个init函数，该函数会在main之前执行，被go的框架调用
2. 通常可以在它里面做i初始化操作

>注意

1. 如果一个文件同时 包含全局变量，init 函数，main 函数，则执行流程是，全局变量->init->main
``` go
package main
import (
	"fmt"
)

var a =test()
func main() {
	fmt.Println("main")
}
func init (){
	fmt.Println("init")
}
func test()int{
	fmt.Println("test")
	return 100
}
```

结果为    
test    
init  
main  
2. init 函数主要的作用，完成一些初始化工作

## 匿名函数
> 一般只使用一次的函数
``` go
package main
import (
	"fmt"
)

func main() {
	a:=func(n1 int,n2 int)int{
		return n1+n2
	}(10,20)
	fmt.Println(a);
}
```
> 全局变量
```go
package main
import (
	"fmt"
)

var (
	F1=func(n1 int ,n2 int) int{
		return n1*n2
	}
)

func main() {
	a:=func(n1 int,n2 int)int{
		return n1+n2
	}(10,20)
	fmt.Println(a);

	fmt.Printf("全局匿名函数 %d",F1(3,4))
}

```
## 闭包
闭包就是一个函数和其他相关的引用环境组合的一个整体  
```go
package main
import (
	"fmt"
)

func main() {
	f:=AddUpper()
	fmt.Println(f(1))  //1
	fmt.Println(f(2))  //13
}
func AddUpper() func(int)int{
	var n int =10
	return func(x int)int{
		n=n+x
		return n
	}
}

```
返回 11  13
>上面的代码拆分就是

```go
package main
import (
	"fmt"
)
func main() {
	var n int=10;
	f:=func(x int)int{  //f:=AddUpper()
		n=n+x
		return n
	}
	fmt.Println(f(1))  //11
	fmt.Println(f(2))  //13
}
```

>案例
根据传入的文件名称，如果没后缀就加上，有就返回
``` go
package main
import (
	"fmt"
	"strings"
)
func main() {
	f:=makeSuffix(".md")
	fmt.Println("文件是",f("123"))
	fmt.Println("文件是",f("123.md"))
}

func makeSuffix(suffix string )func (string) string{
	//suffix  外部传入的 和在这里生命一个差不多
	return func(name string) string{
		//if name 没有后缀就加上，否则就直接返回
		if !strings.HasSuffix(name,suffix) {
			return name+suffix
		}
		return name
	}
}

```
说白了，闭包就是可以保存上一次的引用  

## defer 函数
在函数中，我们经常需要创建资源 比如（数据库连接，文件句柄，锁等）为了在函数执行完毕后，能及时释放资源，go的设计者 提出了defer （延时机制）
```go
package main
import (
	"fmt"
)
func main() {
	res := sum(20,10)
	fmt.Println("res=",res)  //4  32
}

func sum(n1 int,n2 int)int{
	// 当执行defer的时候，暂时不执行，会将defer后面的语句进行压栈，当函数执行完毕后，在从derfer按先入后出的方式出栈执行	
	// 在defer 将语句放入到栈时，也会将相关的值拷贝同时进入栈，
	// 资源释放一般都会这弄
	defer fmt.Println("ok1",n1) //3
	defer fmt.Println("ok2",n2)//2
	n1++  //21
	n2++  //11

	res:=n1+n2  // 32
	fmt.Println("ok3 res=",res)// 1，32
	return res
}

```
结果  
ok3 res= 32  
ok2 10  
ok1 20  
res= 32  



## 函数递归
>实例代码
```go
package main
import (
	"fmt"
)

func main() {
	//A(4)  // 递归方式，
	A4(4);  // 模仿递归调用的流程
}

func A(a int ){
	if(a>2){
		a--
		A(a)
	}
	fmt.Println("a=",a);
}

func A4(a int){
	if(a>2){
		a--
		A3(a)
	}
	fmt.Println("a3=",a);

}

func A3(a int){
	if(a>2){
		a--
		A2(a)
	}
	fmt.Println("a2=",a);

}

func A2(a int){
	//a=2
	if(a>2){
		a--
		A3(a)
	}
	fmt.Println("a1=",a);

}


```
> 递归调用其实表面就是自己调用自己，自己是没法调用自己的，就像自己删除自己没法实现的。程序是这样实现的。自己调用自己的时候，其实就是将自己拷贝了一份，然后再调用如代码  
> 1. 先传入4，条件满足，-- 变成了3 又去调用自己(复制自己)，此时等待3的结果
> 2. 进入到3，条件满足  -- 变成了2 又去调用自己(复制自己)，此时等待2的结果
> 3. 进入到2，条件不满足  --直接打印 出2  返回给2等待地方
> 4. 3接收到了2的结果后执行完if 又执行了打印，2
> 5. 2接受3的结果  执行if  打印 3
> 6. main 接受了4的结果


## 内置函数
1. len 用来求长度 stirng array slice map channel
2. new 用来分配内存，主要用来分配值类型，int float32 返回的指针  
   使用方式   n :=new(int)   值是0

3. make 用来分配内存，主要用来分配引用类型，chan map slice 

# 时间
## 常用格式
``` go
package main
import (
	"fmt"
	"time"
)
func main() {
	now := time.Now() //获取当前时间
	fmt.Printf("时间:%v\n", now)
	year := now.Year()     //年
    month := now.Month()   //月
    day := now.Day()       //日
    hour := now.Hour()     //小时
    minute := now.Minute() //分钟
    second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timestamp1 := now.Unix()     //时间戳
    timestamp2 := now.UnixNano() //纳秒时间戳
    fmt.Printf("现在的时间戳：%v\n", timestamp1)
    fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)


	timeObj := time.Unix(timestamp1, 0) //将时间戳转为时间格式   2023-04-11 19:45:17 +0800 CST

	fmt.Printf("时间戳转换：%v\n", timeObj)

	fmt.Printf("星期：%v\n", now.Weekday().String())
	fmt.Printf("格式指定格式 %v \n",now.Format("2006-01-02 15:04:05"))

}

```
## 常用操作
``` go
package main
import (
	"fmt"
	"time"
)
func main() {
	now := time.Now() //获取当前时间
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println("当前时间增加一小时",later.Format("2006-01-02 15:04:05"))
	fmt.Println("时间差",later.Sub(now))

	//Equal 函数会考虑时区的影响，因此不同时区标准的时间也可以正确比较，Equal 方法和用 t==u 不同，Equal 方法还会比较地点和时区信息。
	fmt.Println("当前时间比较:",now.Equal(now))
	fmt.Println("下一个小时时间比较:",now.Equal(later))
	//判断一个时间点是否在另一个时间点之后：
	fmt.Println("当前时间大于某时间",now.After(now.Add(-time.Hour)))
	//判断一个时间点是否在另一个时间点之前：
	fmt.Println("当前时间小于某时间",now.Before(now.Add(time.Hour)))
}

```
## 定时器
``` go
package main
import (
	"fmt"
	"time"
)
func main() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
    for i := range ticker {
        fmt.Println(i) //每秒都会执行的任务
    }

}
```

## 转换
```go
package main
import (
	"fmt"
	"time"
)
func main() {
	var layout string = "2006-01-02 15:04:05"
    var timeStr string = "2023-04-11 15:22:12"
    timeObj1, _ := time.Parse(layout, timeStr)
    fmt.Println(timeObj1)
    timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
    fmt.Println(timeObj2)

}
```



# 包
> go的每一个文件都是一个包，其实就和命名空间一样  
> 用import "包路径"
> 如果提示找不到包 可以关闭mod  go env -w GO111MODULE=off  

![包结构](img/2023-04-09131904.png "包结构")

``` go
package main
import (
	"fmt"
	 "goCode/utils"	//这里直接是文件夹
	 //u "goCode/utils"  //u 是取的别名
)

func main() {

	fmt.Println("333")
	result:=utils.Cal(2.3,2.7,'+')  // 这里是 包名.函数
	fmt.Println("result=",result)
}



package utils
import (
	"fmt"
	
)

//首写比较是大写 才能被外部使用
func Cal(n1 float64,n2 float64,o byte ) float64{
	switch o {
		case '+':
			return n1+n2
		case '-':
			return n1-n2
		case '*':
			return n1*n2
		case '/':
			return n1/n2
		default:
			fmt.Println("不支持的计算")
		return -99999
	}
} 

```



# 编译项目
1. 编译只需要要编译main所在目录
2. go build -o bin\ gocode\.....\main 执行后会生成 exe文件 会放到bin里面
3. pkg  库文件

# 异常处理



默认情况下，程序报错后会直接退出，

我们希望程序报错后，可以捕获异常，并进行处理，保证程序可以正常执行

go语言处理方式有 defer   panic  recover

`defer 在panic之前执行`

## panic

这玩意执行了，是不能恢复的。直接终止了

## defer+recover处理

defer 是函数执行完后执行的， 如果有多个，他会从后往前执行。也就是会先执行最后一个

```go
package main
import (
	"fmt"
)

func test(){
	defer func(){
		err:=recover()  // reconver 系统内置函数，可以捕获异常
		if err!=nil{
			// 抛出异常
			fmt.Println("err=",err)
		}
	}()
	n1:=10
	n2:=0
	n:=n1/n2 // 这里会报错
	fmt.Println("n=",n)
}

func main() {
	test()

	fmt.Println("执行完毕！")
}


```

## 自定义错误

```go
package main
import (
	"fmt"
	"errors"
)

func test(n int)(err error){
	if n==1{
		return nil
	}else{
		//返回自定义错误
		return errors.New("发生了错误")
	}
}

func test1(){
	err:=test(2)
	if err!=nil{
		//如果发生错误，就输出这个错误，并且终止程序
		panic(err)  
	}
	fmt.Println("发生错误后这里是不执行的")
}

func main() {
	test1()

	fmt.Println("执行完毕！")
}
```

# 命令行参数

Variables   

- os.Agrs提供了简单的命令行参数，以命令行参数个数作为标识，参数列表是一个切片，索引0代表程序本身，1代表第一个参数，以此类推，没有更细粒度的参数区分，使用起来简单
- flag提供了更为科学的命令行参数处理办法，提供更细粒度的和更全的参数解析，推荐使用

### os.Agrs

os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。

``` go
package main
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数",len(os.Args))
	for i,v:=range os.Args{
		fmt.Printf("args[%v]=%v\n",i,v)
	}

}

```

> 先执行go build main.go 生成exe程序
>
> cmd 进入  main.exe  参数1 参数2 惨数3

## flag包

> Args 的方式 传参数 必须遵守 顺序，解决这问题可以用flag 包 来解析命令行    参数可以随意

```go
package main
import (
	"fmt"
	"flag"
)

func main() {

	var user string
	var pwd string
	var host string
	var port int

	// &user 就是接用户命令行输入的-u 
	// u    就是-u指定参数
	// ""  默认为空
	//  说明
	flag.StringVar(&user,"u","","用户名，默认为空")
	flag.StringVar(&pwd,"p","","密码，默认为空")
	flag.StringVar(&host,"h","","host,默认为空")
	flag.IntVar(&port,"t",3306,"端口,默认为空")  // intvar

	// 转换
	flag.Parse()
	//main.exe -u root -p 123456 -h 192.168.2.3
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v",user,pwd,host,port)

}


```





