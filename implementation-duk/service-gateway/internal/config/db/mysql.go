package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type MySQL struct {
	db *gorm.DB
}

func NewMySQL(dataSource DatabaseProperty) (DataSource, error) {
	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dataSource.UserName, dataSource.Password, dataSource.Host, dataSource.Port, dataSource.Database)
	db, err := gorm.Open("mysql", dbConnString)
	// set performance related options
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return &MySQL{db}, err
}

func (m *MySQL) GetConnection() *gorm.DB {
	return m.db
}

func (m *MySQL) GetDialect() string {
	return "mysql"
}

func (m *MySQL) OnDebugging() {
	m.db.Debug()
	m.db.LogMode(true)
}