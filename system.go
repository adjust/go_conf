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

func DropPrivileges(user string) {
	//drop privileges if we are on the servers. testing is done on gentoo systems too.
	if os.Getenv("GO_ENV") == "production" || os.Getenv("GO_ENV") == "test" {
		if syscall.Getuid() == 0 {
			uid := 42
			err := syscall.Setuid(uid)
			if err != nil {
				panic(err)
			}
		}
	}
}

func StartSignalCatcher() {
	//react to sighup
	go signalCatcher()
}
