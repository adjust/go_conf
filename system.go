package go_conf

import (
	"os"
	"os/signal"
	"os/user"
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

func DropPrivileges(user_name string) {
	//drop privileges if we are on the servers. testing is done on gentoo systems too.
	if os.Getenv("GO_ENV") == "production" || os.Getenv("GO_ENV") == "test" {
		if syscall.Getuid() == 0 {
			usr, err := user.Lookup(user_name)
			if err != nil {
				panic(err)
			}
			err = syscall.Setuid(usr.Uid)
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
