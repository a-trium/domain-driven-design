package config

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/cart"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/discount"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/order"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/product"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
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
const SQLITE_FILE_NAME = "local.db"
const SQLITE3 = "sqlite3"

func GetDatabase(env *Environment, log *Logger) *DBConnection {

	var db *gorm.DB
	var err error
	var dialect = SQLITE3

	dbProperty := env.DatabaseProperty

	if env.IsProd() {
		dialect = "??"
	} else if env.IsDev() {
		dialect = "mysql"
		db, err = connectMysql(dbProperty)
	} else {
		dialect = SQLITE3
		db, err = openSqlite3(ASSET_DIR_PATH + SQLITE_FILE_NAME)
	}

	if err != nil {
		log.Fatalw("Failed to connect DB", "error", err)
		panic(err)
	}
	log.Infow("Database connected", "dialect", dialect)

	if !env.IsProd() && env.isDebugging() {
		log.Infow("Debugging Mode!")
		db = db.LogMode(true)
		db = db.Debug()
	}

	if isSqLite(dialect) {
		db.Set("gorm:table_options", "").AutoMigrate(
			&user.Customer{},
			&user.Seller{},
			&product.Tag{},
			&product.Product{},
			&product.Option{},
			&order.Order{},
			&order.Detail{},
			&cart.Cart{},
			&discount.BaseCoupon{})

		db.Exec("PRAGMA foreign_keys = ON;")
	}

	return &DBConnection{db}
}

func isSqLite(dialect string) bool {
	return strings.Compare(dialect, SQLITE3) == 0
}

func openSqlite3(filePath string) (*gorm.DB, error) {
	return gorm.Open("sqlite3", filePath)
}

func connectMysql(dataSource DatabaseProperty) (*gorm.DB, error) {
	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dataSource.UserName, dataSource.Password, dataSource.Host, dataSource.Port, dataSource.Database)
	db, err := gorm.Open("mysql", dbConnString)
	// set performance related options
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db, err
}
