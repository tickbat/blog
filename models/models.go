package models

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"blog/pkg/setting"
)

type model struct {
	ID int `gorm:"primary_key json:"id"`
	CreatedOn int `json:"create_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
		dbData = &setting.Config.Database
		dbType = dbData.Type
		dbName = dbData.Name
		user = dbData.Name
		password = dbData.Password
		host = dbData.Host
		tablePrefix = dbData.TablePrefix
	)
	db, err := gorm.Open(dbType, "%s:%s@/tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName
	)
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	    return tablePrefix + defaultTableName;
	}
	db.SingularTable(true)

  	defer db.Close()
}