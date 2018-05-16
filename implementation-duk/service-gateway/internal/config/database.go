package config

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

)

type DatabaseProperty struct {
	Host     string `default:"localhost"`
	Port     string `default:"3306"`
	UserName string `default:"root"`
	Password string `default:"root"`
	Database string `default:"application"`
}

var (
	db *gorm.DB
	err error
)

func GetDatabase(property DatabaseProperty) *gorm.DB {

	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		property.UserName, property.Password, property.Host, property.Port, property.Database)
	db, err = gorm.Open("mysql", dbConnString)

	if err != nil {
		panic(err)
	}

	// migration
	db.SingularTable(true)

	option := "ENGINE=InnoDB"

	// Automigrate
	db.Set("gorm:table_options", option).AutoMigrate(&domain.User{})


	if !IsProdMode() {
		db = db.LogMode(true)
		db = db.Debug()
	}

	return db
}
