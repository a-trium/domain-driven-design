package db

import (
	"github.com/jinzhu/gorm"
)

type SQLite struct {
	db *gorm.DB
}

func NewSQLite(filePath string) (DataSource, error) {
	db, err := gorm.Open("sqlite3", filePath)
	return &SQLite{db}, err
}

func (s *SQLite) GetConnection() *gorm.DB {
	return s.db
}

func (s *SQLite) GetDialect() string {
	return "sqlite3"
}

func (s *SQLite) OnDebugging() {
	s.db.LogMode(true)
	s.db.Debug()
}
