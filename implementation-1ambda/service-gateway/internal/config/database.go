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
		db.LogMode(true)
		db.Debug()
	}

	// migration
	if useSqlite {
		option := ""
		db.Set("gorm:table_options", option).AutoMigrate(
			&user.User{},
			&user.AuthIdentity{},
			&product.Category{},
			&product.Image{},
			&product.Product{},
			&order.Order{},
			&order.OrderDetail{},
		)

		// SQLite doesn't support `ADD CONSTRAINT`
		// - https://github.com/jinzhu/gorm/blob/b2b568daa8e27966c39c942e5aefc74bcc8af88d/association_test.go#L846
		// db.Model(&user.AuthIdentity{}).AddForeignKey("user_id", "User(id)", "RESTRICT", "CASCADE")

		// Foreign key constraint is disabled by default in SQLite for backward compatibility
		// - http://sqlite.org/foreignkeys.html
		db.Exec("PRAGMA foreign_keys = ON;")

	} else {
		Migrate(db.DB(), dialect)
	}

	return db
}

func getDbLogger() *zap.SugaredLogger {
	logger := GetLogger().With("context", "database")

	return logger
}

func Migrate(db *sql.DB, dialect string) {
	logger := getDbLogger()

	migrationSrc := &migrate.PackrMigrationSource{
		Box: packr.NewBox(SqlSchemaDir),
	}

	migrations, err := migrationSrc.FindMigrations()
	if err != nil {
		logger.Fatalw("Failed to find sql migrations")
	}

	appliedMigrationCount, err := migrate.Exec(
		db,
		dialect,
		migrate.MemoryMigrationSource{Migrations: migrations},
		migrate.Up,
	)

	if err != nil {
		failedMigr := migrations[appliedMigrationCount]

		logger.Warnw("Found sql migration error. Doing rollback...", "down", failedMigr.Down)
		_, downErr := migrate.Exec(
			db,
			dialect,
			migrate.MemoryMigrationSource{Migrations: []*migrate.Migration{failedMigr}},
			migrate.Down,
		)

		if downErr != nil {
			logger.Errorw("Failed to do rollback sql migration", "error", downErr)
		}

		logger.Fatalw("Failed to do sql migration", "error", err)
	}

	totalMigrationCount := len(migrations)
	if appliedMigrationCount != totalMigrationCount {
		logger.Infow("Some migrations are skipped", "total", totalMigrationCount, "applied", appliedMigrationCount)
	}

	for i := 0; i < totalMigrationCount; i++ {
		skipped := true

		if totalMigrationCount-appliedMigrationCount <= i {
			skipped = false
		}

		logger.Infow("Migration File", "filename", migrations[i].Id, "skip", skipped)
	}

	logger.Infow("Finished migration")
}
