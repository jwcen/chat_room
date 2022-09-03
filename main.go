package main

import (
	"os"
	"os/signal"
	"syscall"
	
	"github.com/zserge/lorca"
)

func main() {
	var ui lorca.UI
	ui, _ = lorca.New("https://www.baidu.com", "", 800, 600, "--disable-sync", "--disable-translate")

	/*
		关闭UI后，主线程会自动退出
		中断主线程后，UI会自动退出
	*/
	chSignal := make(chan os.Signal, 1) // 用来接收操作系统的信号
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)  // 订阅SIGINT(中断)，SIGTERM（终止）信号

	// select会 等待 第一个可以读或写的channel进行操作
	select { // 阻塞当前线程（关闭ui或者中断），直到其中一个有值
	case <-ui.Done():
	case <-chSignal:
	}
	ui.Close()
	// select {
	// case <-ui.Done():
	// 	ui.Close()
	// 	return
	// case <-chSignal:
	// 	ui.Close()
	// 	return
	// }
	// return
}