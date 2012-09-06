package go_conf

import (
	"os"
	"os/signal"
	"os/user"
	"strconv"
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
	//drop privileges only if we are on the servers.
	if os.Getenv("GO_ENV") == "production" {
		if syscall.Getuid() == 0 {
			usr, err := user.Lookup(user_name)
			if err != nil {
				panic(err)
			}
			id, _ := strconv.Atoi(usr.Uid)
			err = syscall.Setuid(id)
			if err != nil {
				panic(err)
			}
			logger.Println("Dropped Privileges to " + usr.Uid)
		}
	}
}

func StartSignalCatcher() {
	//react to sighup
	go signalCatcher()
}
