package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL MySQL  `yaml:"mysql"`
	Lock  string `yaml:"lock"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var AppConfig Config

var SqlDB *sql.DB

func ReadConfig() error {
	configYaml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println("Config Error")
		return err
	}

	var config Config

	err = yaml.Unmarshal(configYaml, &config)

	if err != nil {
		log.Println("Config unmarshal error")
	}
	// log.Printf("%+v\n", config)
	AppConfig = config
	return err
}

func GetConfig() Config {
	return AppConfig
}

func GetSqlDB() *sql.DB {
	return SqlDB
}

func InitMysqlConnection() error {
	config := GetConfig()
	uri := fmt.Sprintf("%s:%s@tcp(%s)", config.MySQL.Username, config.MySQL.Password, config.MySQL.Host)
	log.Println("Connecting to " + uri)
	db, err := sql.Open("mysql", uri+"/hackernews")
	if err != nil {
		return err
	}

	log.Println("Connected to " + uri)

	SqlDB = db

	return nil
}
