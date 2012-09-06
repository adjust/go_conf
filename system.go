package go_conf

import (
	"os"
	"os/signal"
	"syscall"
)

func signalCatcher() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	for signal := range ch {
		if signal == syscall.SIGHUP {
			logger.Println("received SIGHUP exiting...")
			os.Exit(0)
		}
	}
}

func StartSignalCatcher() {
	//react to sighup
	go signalCatcher()
}
