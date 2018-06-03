package config

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config/db"
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

const ASSET_DIR_PATH = "implementation-duk/service-gateway/asset/"
const SQLITE_FILE_NAME = "local.db"

func GetDatabase(env *Environment, log *Logger) db.DataSource {

	var dataSource db.DataSource
	var err error

	dbProperty := env.DatabaseProperty

	if env.IsProd() {

	} else if env.IsDev() {
		dataSource, err = db.NewMySQL(dbProperty)
	} else {
		dataSource, err = db.NewSQLite(ASSET_DIR_PATH + SQLITE_FILE_NAME)
	}

	if err != nil {
		log.Fatalw("Failed to connect DB", "error", err)
		panic(err)
	}
	log.Infow("Database connected", "dialect", dataSource.GetDialect())

	if !env.IsProd() && env.IsDebugging() {
		log.Infow("Debugging Mode!")
		dataSource.OnDebugging()
	}

	if isLocalMemoryDB(env, dataSource) {
		onAutoMigration(dataSource.GetConnection())
	}

	return dataSource
}

func isLocalMemoryDB(env *Environment, dataSource db.DataSource) bool {
	return env.IsLocal() && isSQLite3(dataSource.GetDialect())
}

func isSQLite3(dialect string) bool {
	return strings.EqualFold(dialect,"sqlite3")
}


func onAutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "").AutoMigrate(
		&user.Customer{},
		&user.Seller{},
		&product.Product{},
		&product.Option{},
		&product.Tag{},
		&order.Order{},
		&order.Detail{},
		&cart.Cart{},
		&discount.BaseCoupon{})

	db.Exec("PRAGMA foreign_keys = ON;")
}
