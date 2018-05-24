package config

import (
	"database/sql"
	"fmt"
	"os/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rubenv/sql-migrate"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

const SqlSchemaDir = "../../migrations/"

var (
	db *gorm.DB
	err error
)

func GetDatabase() *gorm.DB {

	logger := getDbLogger()

	dialect := "mysql"
	useSqlite := Env.IsTestMode()

	// Use sqlite3 for `TEST` env
	if useSqlite {
		dialect = "sqlite3"
		uuidString := uuid.NewV4().String()
		filename := fmt.Sprintf("/tmp/ddd_gateway_%s.db", uuidString)
		db, err = gorm.Open(dialect, filename)
	} else {
		dialect = "mysql"
		dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			Env.MysqlUserName, Env.MysqlPassword, Env.MysqlHost, Env.MysqlPort, Env.MysqlDatabase)
		db, err = gorm.Open(dialect, dbConnString)

		// set performance related options
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)

		logger.Infow("MySQL Connection", "dbConnString", dbConnString)
	}

	// check connection
	if err != nil {
		logger.Fatalw("Failed to connect DB", "error", err)
	}
	logger.Infow("Database connected", "dialect", dialect)

	// set gorm options
	db.SingularTable(true)
	if (Env.IsLocalMode() && Env.DebugEnabled()) || Env.IsTestMode() {
		db = db.LogMode(true)
		db = db.Debug()
	}

	// migration
	if useSqlite {
		option := ""
		db.Set("gorm:table_options", option).AutoMigrate(
			&user.User{},
		)

		//db.Model(&product.Product{}).AddForeignKey("category_id", "Category(id)", "RESTRICT", "CASCADE")

	} else {
		doMigration(db.DB(), dialect)
	}

	if !IsProdMode() {
		db = db.LogMode(true)
		db = db.Debug()
	}

	return db
}


func getDbLogger() *zap.SugaredLogger {
	logger := GetLogger().With("context", "db")

	return logger
}

func doMigration(db *sql.DB, dialect string) {
	logger := getDbLogger()
	logger.Infow("starting Migration", "SqlSchemaDir", SqlSchemaDir)

	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox(SqlSchemaDir),
	}

	findedMigrations, err := migrations.FindMigrations()
	if err != nil {
		logger.Fatalw("Failed to find sql migrations")
	}

	for _, migr := range findedMigrations {
		m := []*migrate.Migration{migr}
		n, err := migrate.Exec(
			db,
			dialect,
			migrate.MemoryMigrationSource{Migrations: m,},
			migrate.Up,
		)

		if err != nil {
			logger.Warnw("Found sql migration error. Doing rollback...", "down", migr.Down)
			_, downErr := migrate.Exec(
				db,
				dialect,
				migrate.MemoryMigrationSource{Migrations: []*migrate.Migration{migr}},
				migrate.Down,
			)

			if downErr != nil {
				logger.Errorw("Failed to do rollback sql migration", "error", downErr)
			}

			logger.Fatalw("Failed to do sql migration", "error", err)
		}

		if n != 0 {
			logger.Infow("Finished migration", "file", migr.Id)
		} else {
			logger.Infow("Skip migration", "file", migr.Id)
		}
	}
}