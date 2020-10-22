package main

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	Config *DataBases
)

type DataBases struct {
	Database []string `yaml:"Database"`
	Address  string   `yaml:"Address"`
	Username string   `yaml:"Username"`
	Password string   `yaml:"Password"`
}

func init() {
	InitConfig()
}

func InitConfig() {
	LoadConfig("app.yaml")
}

// LoadConfig ...
func LoadConfig(file string) {
	if Config != nil {
		return
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return
	}

	Config = new(DataBases)

	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %v", err)
	}

	viper.Unmarshal(&Config)

}

func dbExec(file string) (map[string]error, error) {
	sqlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	failedDB := make(map[string]error)
	for _, database := range Config.Database {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
			Config.Username, Config.Password, Config.Address, database))
		defer db.Close()
		if err != nil {
			if _, ok := failedDB[database]; !ok {
				failedDB[database] = err
			}
			continue
		}
		tx, err := db.Begin()
		defer tx.Rollback()
		if err != nil {
			if _, ok := failedDB[database]; !ok {
				failedDB[database] = err
			}
			break
		}
		execs := strings.Split(string(sqlFile), ";")
		for i := 0; i < len(execs)-1; i++ {
			_, err = tx.Exec(execs[i])
			if err != nil {
				if _, ok := failedDB[database]; !ok {
					failedDB[database] = err
				}
				continue
			}
		}
		tx.Commit()
	}

	return failedDB, nil
}

func main() {
	failedDB, err := dbExec("test.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(failedDB) == 0 {
		fmt.Println("all success")
		return
	}
	for db, failedErr := range failedDB {
		fmt.Printf("DB: %s, error: %v", db, failedErr)
	}
}
