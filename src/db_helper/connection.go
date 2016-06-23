package db_helper

import (
	"sync"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "sync"
	"github.com/jinzhu/gorm"
	"github.com/Sirupsen/logrus"
)

var dbConnection *gorm.DB
var once sync.Once

// Abbr. Shared Connection cool to connect with mysql database
// GoLang DB instance automatically handles database pooling
func SharedConnection() (*gorm.DB) {
	once.Do(func() {
		var err error
		// Setting database endpoint
		dbConnection,err = gorm.Open("mysql", "root:root@tcp(snabar.com:3306)/snabar_staging?parseTime=true")
		//dbConnection, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/snabar_staging")
		dbConnection.LogMode(true)
		if err != nil {
			logrus.Info("Error encoutered while getting connection from mysql")
			fmt.Println(err)
		}
		logrus.Info("Getting connection from MYSQL")
		return
	})
	return dbConnection
}
