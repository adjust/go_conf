package go_conf

import (
	"log"
	"strconv"
)

func GetRedisConf() (redis_host string, redis_db int) {
	redis_host = getConfigParameter("redis", "host")
	db_str := getConfigParameter("redis", "db")
	redis_db, err := strconv.Atoi(db_str)
	if err != nil {
		log.Panic("redis db not an integer!")
	}
	return
}

func GetMasterRedisConf() (redis_host string, redis_db int) {
	redis_host = getConfigParameter("master_redis", "host")
	db_str := getConfigParameter("master_redis", "db")
	redis_db, err := strconv.Atoi(db_str)
	if err != nil {
		log.Panic("master redis db not an integer!")
	}
	return
}
