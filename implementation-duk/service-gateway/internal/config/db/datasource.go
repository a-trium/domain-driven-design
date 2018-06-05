package db

import (
	"github.com/jinzhu/gorm"
)

type DatabaseProperty struct {
	Host     string `default:"localhost"`
	Port     string `default:"3306"`
	UserName string `default:"root"`
	Password string `default:"root"`
	Database string `default:"application"`
}

type DataSource interface {
	GetConnection() *gorm.DB
	GetDialect() string
	OnDebugging()
}