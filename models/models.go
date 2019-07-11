package models

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"blog/pkg/setting"
)

type Model struct {
	ID int `gorm:"primary_key json:"id"`
	CreatedOn int `json:"create_on"`
	ModifiedOn int `json:"modified_on"`
}

var db *gorm.DB

func init() {
	var (
		err error
		dbData = setting.Config.Database
		dbType = dbData.Type
		dbName = dbData.Name
		user = dbData.User
		password = dbData.Password
		host = dbData.Host
		tablePrefix = dbData.TablePrefix
	)
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@/tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	// db, err = gorm.Open(dbType, "root:root@/tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	    return tablePrefix + defaultTableName;
	}
	db.SingularTable(true)

  	defer db.Close()
}

func CloseDB() {
	defer db.Close()
}