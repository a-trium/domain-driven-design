package config

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DatabaseProperty struct {
	Host     string `default:"localhost"`
	Port     string `default:"3306"`
	UserName string `default:"root"`
	Password string `default:"root"`
	Database string `default:"application"`
}

type DBConnection struct {
	db *gorm.DB
}

func (c *DBConnection) GetDB() *gorm.DB {
	return c.db
}

const ASSET_DIR_PATH = "implementation-duk/service-gateway/asset/"
const SQLITE_FILE_NAME = "gateway.db"

func GetDatabase(env *Environment) *DBConnection {

	var db *gorm.DB
	var err error

	dbProperty := env.DatabaseProperty

	if env.IsProd() {

	} else if env.IsDev() {
		dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbProperty.UserName, dbProperty.Password, dbProperty.Host, dbProperty.Port, dbProperty.Database)
		db, err = gorm.Open("mysql", dbConnString)
	} else {
		db, err = gorm.Open("sqlite3", ASSET_DIR_PATH + SQLITE_FILE_NAME)
	}

	if err != nil {
		panic(err)
	}


	db.AutoMigrate(&user.Customer{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user.Customer{})

	//if !IsProdMode() {
	//	db = db.LogMode(true)
	//	db = db.Debug()
	//}

	return &DBConnection{db}
}
