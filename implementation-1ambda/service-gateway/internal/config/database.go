package config

import (
	"database/sql"
	"fmt"

	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/order"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/product"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rubenv/sql-migrate"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

const SqlSchemaDir = "../../asset/sql"

func GetDatabase() *gorm.DB {
	logger := getDbLogger()

	var db *gorm.DB
	var err error

	useSqlite := Env.IsTestMode()
	dialect := "sqlite3"

	// Use sqlite3 for `TEST` env
	if useSqlite {
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
	}

	if err != nil {
		logger.Fatalw("Failed to connect DB", "error", err)
	}
	logger.Infow("Database connected", "dialect", dialect)

	// set gorm options
	db.SingularTable(true)
	if (Env.IsLocalMode() && Env.DebugEnabled()) || (Env.IsTestMode() && Env.DebugEnabled()) {
		db = db.LogMode(true)
		db = db.Debug()
	}

	// migration
	if useSqlite {
		option := ""
		db.Set("gorm:table_options", option).AutoMigrate(
			&user.User{},
			&product.Category{},
			&product.Image{},
			&product.Product{},
			&order.Order{},
			&order.OrderDetail{},
		)

		//db.Model(&product.Product{}).AddForeignKey("category_id", "Category(id)", "RESTRICT", "CASCADE")

	} else {
		doMigration(db.DB(), dialect)
	}

	return db
}

func getDbLogger() *zap.SugaredLogger {
	logger := GetLogger().With("context", "database")

	return logger
}

func doMigration(db *sql.DB, dialect string) {
	logger := getDbLogger()

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
