package config

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

)

func GetDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", "root", "localhost", "3306", "application")
	db, err = gorm.Open("mysql", dbConnString)

	if err != nil {
		panic(err)
	}

	// migration
	db.SingularTable(true)

	option := "ENGINE=InnoDB"

	// Automigrate
	db.Set("gorm:table_options", option).AutoMigrate(&domain.User{})

	//if (Env.IsLocalMode() && Env.DebugEnabled()) || Env.IsTestMode() {
		db = db.LogMode(true)
		db = db.Debug()
	//}

	return db
}
