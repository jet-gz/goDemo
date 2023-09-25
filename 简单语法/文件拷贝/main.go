package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func CopyFile(distName string,srcName string) (wtitten int64,err error){
	srcFile,err:=os.Open(srcName)
	if err!=nil{
		fmt.Println("错误信息",err)
	}
	defer srcFile.Close()
	// 读取 目标
	reader:=bufio.NewReader(srcFile)

	dstFile,err:=os.OpenFile(distName,os.O_WRONLY | os.O_CREATE,0666)
	if err!=nil{
		fmt.Println("错误",err)
		return
	}
	defer dstFile.Close()
	// 获取dst的writer
	writer:=bufio.NewWriter(dstFile)
	return io.Copy(writer,reader)
}

func main() {
	src_path:="F:/software/cn_office_professional_plus_2019_x86_x64_dvd_5e5be643.iso"
	path:="E:/123.iso"
    _ ,err:=	CopyFile(path,src_path)
	if err!=nil{
		fmt.Println("错误",err)
	}else{
		fmt.Println("拷贝完成")
	}
}


