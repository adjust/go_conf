package go_conf

func GetAmqpConf() string {
	amqp_user := getConfigParameter("amqp", "user")
	amqp_pass := getConfigParameter("amqp", "pass")
	ampq_host := getConfigParameter("amqp", "host")
	amqp_port := getConfigParameter("amqp", "port")
	return "amqp://" + amqp_user + ":" + amqp_pass + "@" + ampq_host + ":" + amqp_port + "/"
}
