package go_conf

import (
	"flag"
	"log"
	"os"
)

var (
	port  = flag.String("port", "8080", "which port to listen on? (only applies to servers)")
	shard = flag.Int64("shard", 0, "the id of this proxy (used for sharding)")
)

func GetPort() string {
	return *port
}

func GetShard() int64 {
	if *shard == 0 {
		log.Panic("no valid shard id!")
		os.Exit(1)
	}
	return *shard
}
