package config

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config/db"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/cart"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/discount"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/order"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/product"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/tag"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/gobuffalo/packr"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rubenv/sql-migrate"
	"strings"
)

const ASSET_RELATIVE_PATH = "../../asset/data"
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
		flyway(log, dataSource)
	}

	return dataSource
}

func isLocalMemoryDB(env *Environment, dataSource db.DataSource) bool {
	return env.IsLocal() && isSQLite3(dataSource.GetDialect())
}

func isSQLite3(dialect string) bool {
	return strings.EqualFold(dialect, "sqlite3")
}

func onAutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "").AutoMigrate(
		&user.Customer{},
		&user.Seller{},

		&product.Product{},
		&product.Option{},
		&product.ProductTag{},

		&tag.Tag{},

		&order.Order{},
		&order.Detail{},
		&cart.Cart{},
		&discount.Coupon{})

	db.Exec("PRAGMA foreign_keys = ON;")
}

func flyway(log *Logger, dataSource db.DataSource) {
	migrationSrc := &migrate.PackrMigrationSource{
		Box: packr.NewBox(ASSET_RELATIVE_PATH),
	}

	migrations, err := migrationSrc.FindMigrations()
	if err != nil {
		log.Fatalw("Failed to find sql migrations", err)
	}

	appliedMigrationCount, err := migrate.Exec(
		dataSource.GetConnection().DB(),
		dataSource.GetDialect(),
		migrate.MemoryMigrationSource{Migrations: migrations},
		migrate.Up,
	)

	if err != nil {
		failedMigr := migrations[appliedMigrationCount]

		log.Warnw("Found sql migration error. Doing rollback...", "down", failedMigr.Down)
		_, downErr := migrate.Exec(
			dataSource.GetConnection().DB(),
			dataSource.GetDialect(),
			migrate.MemoryMigrationSource{Migrations: []*migrate.Migration{failedMigr}},
			migrate.Down,
		)

		if downErr != nil {
			log.Errorw("Failed to do rollback sql migration", "error", downErr)
		}

		log.Fatalw("Failed to do sql migration", "error", err)
	}

	totalMigrationCount := len(migrations)
	if appliedMigrationCount != totalMigrationCount {
		log.Infow("Some migrations are skipped", "total", totalMigrationCount, "applied", appliedMigrationCount)
	}

	for i := 0; i < totalMigrationCount; i++ {
		skipped := true

		if totalMigrationCount-appliedMigrationCount <= i {
			skipped = false
		}

		log.Infow("Migration File", "filename", migrations[i].Id, "skip", skipped)
	}
	log.Infow("Finished migration")
}
