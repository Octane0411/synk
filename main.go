package main

import (
	"github.com/Octane0411/synk/server"
	"github.com/zserge/lorca"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
) 

func main() {
	go server.Run()

	ui, _ := lorca.New("http://127.0.0.1:27149/static/index.html", "", 800, 600, "--disable-sync", "--disable-translate")
	//cmd := startBrowser()
	chSignal := listenToInterrupt()
	select {
	case <-chSignal:
		ui.Close()
	case <-ui.Done():
		os.Exit(0)
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

func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	return chSignal
}
