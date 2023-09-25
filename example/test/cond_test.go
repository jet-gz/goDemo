package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func listen(x int) {
	// 获取锁
	cond.L.Lock()
	// 等待通知  暂时阻塞
	cond.Wait()
	fmt.Println(x)
	// 释放锁
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 60; i++ {
		go listen(i)
	}
	fmt.Println("start all")

	fmt.Println("+++++++++++++++++获取一个")
	cond.Signal()

	// 3秒之后 下发一个通知给已经获取锁的goroutine
	fmt.Println("++++++++++++++++++++3s 后唤醒一个")
	time.Sleep(time.Second * 3)
	cond.Signal()

	// 3秒之后 下发广播给所有等待的goroutine
	fmt.Println("++++++++++++++++++++3s后唤醒所有的")
	time.Sleep(time.Second * 3)

	cond.Broadcast()
	// 阻塞直到所有的全部输出
	time.Sleep(time.Second * 4)

}
