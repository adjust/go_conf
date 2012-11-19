package go_conf

import "flag"

var (
	port  = flag.String("port", "8080", "which port to listen on? (only applies to servers)")
	shard = flag.Int64("shard", 1, "the id of this proxy (used for sharding)")
)

func GetPort() string {
	return *port
}

func GetShard() int64 {
	return *shard
}
