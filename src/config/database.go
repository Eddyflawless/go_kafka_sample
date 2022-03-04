package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration

type DBConfig struct {
	Host string
	Port string
	User string
	DBName string
	Password string
}

func BuildConfig(host string, port string, user string, dbName string, password string) *DBConfig {

	dbConfig := DBConfig{
		Host: host,
		Port: port,
		User: user,
		DBName: dbName,
		Password: password,
	}

	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

}