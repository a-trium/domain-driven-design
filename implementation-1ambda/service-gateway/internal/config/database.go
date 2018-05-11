package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
)

func GetDatabase() *gorm.DB {
	logger := GetLogger()

	var db *gorm.DB
	var err error

	// Use sqlite3 for `TEST` env
	if Env.IsTestMode() {
		uuidString := uuid.NewV4().String()
		filename := fmt.Sprintf("/tmp/go-ref_gateway_%s.db", uuidString)
		logger.Info("Use sqlite3 database")
		db, err = gorm.Open("sqlite3", filename)
	} else {
		logger.Info("Use mysql database")
		dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			Env.MysqlUserName, Env.MysqlPassword, Env.MysqlHost, Env.MysqlPort, Env.MysqlDatabase)
		db, err = gorm.Open("mysql", dbConnString)
	}

	if err != nil {
		logger.Fatalw("Failed to connect DB", "error", err)
	}

	// migration
	logger.Info("Migrating tables")
	db.SingularTable(true)

	//option := "ENGINE=InnoDB"
	//if Env.IsTestMode() {
	//	option = ""
	//}

	// Automigrate
	// db.Set("gorm:table_options", option).AutoMigrate(&Session{})

	if !Env.IsTestMode() {
		// https://github.com/jinzhu/gorm/issues/1824#issuecomment-378123682
		// gorm doesn't generate FK w/ `AutoMigrate`

		// db.Model(&BrowserHistory{}).AddForeignKey("session_id", "session(session_id)", "RESTRICT", "CASCADE")
	}

	if (Env.IsLocalMode() && Env.DebugEnabled()) || Env.IsTestMode() {
		db = db.LogMode(true)
		db = db.Debug()
	}

	return db
}
