package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var brands = []string{"lv", "ls", "3h", "bh", "c7", "c8", "cdd", "co", "dm", "hb", "xunya", "qm", "sc", "tz"}

func main() {

}

type Count struct {
	Count int `gorm:"column"`
}

var DB *gorm.DB

type MySQL struct {
	DB *gorm.DB
}

// DataBaseConfig ...
type DataBaseConfig struct {
	Username string
	Password string
	Address  string
	Database string
}

// NewMySQL ...
func NewMySQL(dbname string) (*MySQL, error) {
	db := new(MySQL)
	err := db.Connect(dbname)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Connect ...
func (db *MySQL) Connect(dbname string) error {
	var err error
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"rd_report",
		"bkreport@DEV",
		"10.200.252.216:3306",
		dbname,
	)

	db.DB, err = gorm.Open("mysql", connect)
	if err != nil {
		return err
	}

	return nil
}

// Session ...
func (db *MySQL) Session() *gorm.DB {
	return db.DB
}

// Begin ...
func (db *MySQL) Begin() *gorm.DB {
	return db.DB.Begin()
}

// Close ...
func (db *MySQL) Close() {
	db.DB.Close()
}
