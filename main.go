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
	go server.Run()

	//var ui lorca.UI
	//ui, _ = lorca.New("https://baidu.com", "", 800, 600, "--disable-sync", "--disable-translate")

	cmd := startBrowser()
	chSignal := listenToInterrupt(cmd)
	select {
	case <-chSignal:
		cmd.Process.Kill()
	}

}

func startBrowser() *exec.Cmd {
	chromePath := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	cmd := exec.Command(chromePath, "http://127.0.0.1:27149/static/index.html")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

func listenToInterrupt(cmd *exec.Cmd) chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	return chSignal
}
