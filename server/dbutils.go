package server

import (
	"fmt"
	"log"
	"os"
)

// DatabaseConfig is to collection of properties requires to connect Database
type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	DBDriver string
	DBPort   string
}

func (d *DatabaseConfig) InitializeConfig() {

	var exist bool

	if d.Host, exist = os.LookupEnv(dbhost); !exist {
		log.Fatalf("env %s not found", dbhost)
	}
	if d.User, exist = os.LookupEnv(dbuser); !exist {
		log.Fatalf("env %s not found", dbuser)
	}
	if d.Password, exist = os.LookupEnv(dbpassword); !exist {
		log.Fatalf("env %s not found", dbpassword)
	}
	if d.DBName, exist = os.LookupEnv(dbname); !exist {
		log.Fatalf("env %s not found", dbname)
	}
	if d.DBDriver, exist = os.LookupEnv(dbdriver); !exist {
		log.Fatalf("env %s not found", dbdriver)
	}
	if d.DBPort, exist = os.LookupEnv(dbport); !exist {
		log.Fatalf("env %s not found", dbport)
	}
}

func (d *DatabaseConfig) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", d.User,
		d.Password, d.Host, d.DBPort, d.DBName)
}

