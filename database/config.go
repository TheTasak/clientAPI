package database

import "fmt"

type Config struct {
	ServerName string
	User       string
	Password   string
	Database   string
}

func (config Config) GetConnectionString() string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.Database)
	return connectionString
}
