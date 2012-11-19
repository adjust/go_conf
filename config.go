package go_conf

import (
	"flag"
	"github.com/kylelemons/go-gypsy/yaml"
	"io/ioutil"
	"log"
	"os"
)

var (
	exitHandler   ExitHandler
	config        *yaml.File
	environment   string
	config_file   = flag.String("config", "./config/database.yml", "the database.yml")
	log_file_name = flag.String("log", "./log/server.log", "where does the log go?")
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
