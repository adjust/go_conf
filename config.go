package go_conf

import (
	"flag"
	"github.com/kylelemons/go-gypsy/yaml"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var (
	exitHandler   ExitHandler
	config        *yaml.File
	environment   string
	config_file   = flag.String("config", "./config/database.yml", "the database.yml")
	log_file_name = flag.String("log", "./log/server.log", "where does the log go?")
	port          = flag.String("port", "8080", "which port to listen on? (only applies to servers)")
	shard         = flag.Int64("shard", 0, "the id of this proxy (used for sharding)")
)

func init() {
	exitHandler = &StandardHandler{}
	flag.Parse()
	setEnv()
	initlogAndConfig()
	startSignalCatcher()
}

func initlogAndConfig() {
	//create log   
	log_file, err := os.OpenFile(*log_file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("cannot write log")
	}
	log.SetOutput(log_file)
	log.SetFlags(5)

	//read the config and build config stuff
	c_file, err := ioutil.ReadFile(*config_file)
	if err != nil {
		log.Panic("no config file found")
	}
	config = yaml.Config(string(c_file))
}

func GetEnv() string {
	return environment
}

func setEnv() {
	environment = os.Getenv("GO_ENV")
	if environment == "" {
		environment = "development"
	}
}

func getConfigParameter(prefix, name string) string {
	param, err := config.Get(prefix + "_" + environment + "." + name)
	if err != nil {
		log.Panic("missing config parameter: " + prefix + " " + name)
	}
	return param
}

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

func GetPgConf() string {
	pg_user := getConfigParameter("postgres", "user")
	pg_db := getConfigParameter("postgres", "db")
	pg_host := getConfigParameter("postgres", "host")
	return "user=" + pg_user + " dbname=" + pg_db + " sslmode=disable host=" + pg_host
}

func GetAmqpConf() string {
	amqp_user := getConfigParameter("amqp", "user")
	amqp_pass := getConfigParameter("amqp", "pass")
	ampq_host := getConfigParameter("amqp", "host")
	amqp_port := getConfigParameter("amqp", "port")
	return "amqp://" + amqp_user + ":" + amqp_pass + "@" + ampq_host + ":" + amqp_port + "/"
}

func GetPort() string {
	return *port
}

func GetShard() int64 {
	if *shard == 0 {
		log.Panic("no valid shard id set!")
		os.Exit(1)
	}
	return *shard
}
