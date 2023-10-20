package config

import "fmt"

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
		c.Port,
	)
}
