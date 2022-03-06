package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/guanweiqin/synk/server"
	"github.com/guanweiqin/synk/server/config"
	"github.com/lxn/walk"
)

func startBrower(chChromeDie chan struct{}, chBackendDie chan struct{}) {
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+config.GetPort()+"/static/index.html")
	// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 隐藏调用的外部程序的cmd窗口
	err := cmd.Start()
	if err != nil {
		walk.MsgBox(nil, "提示", err.Error(), walk.MsgBoxOK|walk.MsgBoxIconError)
	}
	go func() {
		<-chBackendDie
		cmd.Process.Kill()
	}()
	go func() {
		cmd.Wait()
		chChromeDie <- struct{}{}
	}()
}

func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}

func main() {
	chChromeDie := make(chan struct{})
	chBackendDie := make(chan struct{})
	go server.Run()
	go startBrower(chChromeDie, chBackendDie)
	chSignal := listenToInterrupt()
	select {
	case <-chSignal:
		chBackendDie <- struct{}{}
	case <-chChromeDie:
		os.Exit(0)
	}
}
