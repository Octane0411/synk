package main

import (
	"github.com/Octane0411/synk/server"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
) //勾

func main() {
	//启动gin服务
	go server.Run()

	//var ui lorca.UI
	//ui, _ = lorca.New("https://baidu.com", "", 800, 600, "--disable-sync", "--disable-translate")

	//先写死路径，后面再改
	//启动chrome
	chromePath := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	cmd := exec.Command(chromePath, "http://127.0.0.1:27149/static/index.html")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	//监听中断信号
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-chSignal:
		cmd.Process.Kill()
	}

}
