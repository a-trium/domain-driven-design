package config

import (
	"fmt"

	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/user/domain"
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

	useSqlite := Env.IsTestMode()

	// Use sqlite3 for `TEST` env
	if useSqlite {
		uuidString := uuid.NewV4().String()
		filename := fmt.Sprintf("/tmp/a-trium/domain-driven-design/service-gateway/%s.db", uuidString)
		logger.Info("Use sqlite3 database")
		db, err = gorm.Open("sqlite3", filename)
	} else {
		logger.Info("Use mysql database")
		dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			Env.MysqlUserName, Env.MysqlPassword, Env.MysqlHost, Env.MysqlPort, Env.MysqlDatabase)
		db, err = gorm.Open("mysql", dbConnString)

		// set performance related options
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
	}

	if err != nil {
		logger.Fatalw("Failed to connect DB", "error", err)
	}

	logger.Info("Migrating tables")
	db.SingularTable(true)

	option := "ENGINE=InnoDB"
	if useSqlite {
		option = ""
	}

	db.Set("gorm:table_options", option).AutoMigrate(&user.User{})

	if !useSqlite {
		// https://github.com/jinzhu/gorm/issues/1824#issuecomment-378123682
		// gorm doesn't generate FK w/ `AutoMigrate`

		//db.Model(&user.User{}).AddForeignKey("session_id", "session(session_id)", "RESTRICT", "CASCADE")
	}

	if (Env.IsLocalMode() && Env.DebugEnabled()) || Env.IsTestMode() {
		db = db.LogMode(true)
		db = db.Debug()
	}

	return db
}
