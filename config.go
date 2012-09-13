package go_conf

import (
	"github.com/droundy/goopt"
	"github.com/kylelemons/go-gypsy/yaml"
	"io/ioutil"
	"log"
	"os"
)

var (
	logger      *log.Logger
	config      *yaml.File
	port        *string
	environment string
)

func InitLoggerAndConfig() *log.Logger {
	//get flags
	config_file := goopt.String([]string{"-c", "--config"}, "./config/database.yml", "the database.yml")
	log_file_flag := goopt.String([]string{"-l", "--log"}, "./log/server.log", "where does the log go?")
	port = goopt.String([]string{"-p", "--port"}, "8080", "which port to listen on? (only applies to servers)")
	goopt.Summary = "a go daemon"
	goopt.Parse(nil)

	//create logger
	log_file, err := os.OpenFile(*log_file_flag, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("cannot write log")
	}
	logger = log.New(log_file, "", 5)

	//read the config and build config stuff
	c_file, err := ioutil.ReadFile(*config_file)
	if err != nil {
		logger.Panic("no config file found")
	}
	config = yaml.Config(string(c_file))

	environment = os.Getenv("GO_ENV")
	if environment == "" {
		environment = "development"
	}
	return logger
}

func GetRedisConf() (redis_host string, redis_db string) {
	var err error
	redis_host, err = config.Get("redis_" + environment + ".host")
	if err != nil {
		logger.Panic("missing config parameter: redis host")
	}
	redis_db, err = config.Get("redis_" + environment + ".db")
	if err != nil {
		logger.Panic("missing config parameter: redis db")
	}
	return
}

func GetPgConf() (pg_conf string) {

	pg_user, err := config.Get("postgres_" + environment + ".user")
	if err != nil {
		logger.Panic("missing config parameter: postgres user")
	}

	pg_db, err := config.Get("postgres_" + environment + ".db")
	if err != nil {
		logger.Panic("missing config parameter: postgres db")
	}

	pg_host, err := config.Get("postgres_" + environment + ".host")
	if err != nil {
		logger.Panic("missing config parameter: postgres host")
	}

	pg_conf = "user=" + pg_user + " dbname=" + pg_db + " sslmode=disable host=" + pg_host
	return
}

func GetAmqpConf() (amqp_conf string) {
	var err error
	amqp_conf, err = config.Get("amqp_" + environment + ".url")
	if err != nil {
		logger.Panic("missing config parameter: amqp url")
	}
	return
}

func GetPort() string {
	return *port
}
