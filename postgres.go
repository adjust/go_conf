package go_conf

func GetPgConf() string {
	pg_user := getConfigParameter("postgres", "user")
	pg_db := getConfigParameter("postgres", "db")
	pg_host := getConfigParameter("postgres", "host")
	return "user=" + pg_user + " dbname=" + pg_db + " sslmode=disable host=" + pg_host
}
