package db

type DbConfig struct {
	Host string
	Prot int
	DbName string
}

func GetDbConfig(host string, port int, dbName string) *DbConfig {
	return &DbConfig{host, port, dbName}
}