package models

type Config struct {
	Server struct {
		Debug     bool
		Port      int
		UseTestDB bool
	}
	DB struct {
		Hostname string
		Port     int
		Username string
		Password string
		DBName   string
	}
	TestDB struct {
		Path string
	}
}
