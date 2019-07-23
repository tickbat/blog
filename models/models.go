package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"blog/pkg/setting"
	"blog/pkg/logging"
	"time"
)

type Model struct {
	ID *int `gorm:"primary_key" json:"id"`
	CreatedOn *int `json:"create_on"`
	ModifiedOn *int `json:"modified_on"`
}

var db *gorm.DB

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
    if !scope.HasError() {
        nowTime := time.Now().Unix()
        if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
            if createTimeField.IsBlank {
                createTimeField.Set(nowTime)
            }
        }

        if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
            if modifyTimeField.IsBlank {
                modifyTimeField.Set(nowTime)
            }
        }
    }
}


func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
    if _, ok := scope.Get("gorm:update_column"); !ok {
        scope.SetColumn("ModifiedOn", time.Now().Unix())
    }
}

func init() {
	var (
		err error
		dbType = setting.Database.Type
		dbName = setting.Database.Name
		user = setting.Database.User
		password = setting.Database.Password
		host = setting.Database.Host
		tablePrefix = setting.Database.TablePrefix
	)
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)
	db, err = gorm.Open(dbType, source)
	if err != nil {
		logging.Info("open mysql error: ", err.Error())
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	    return tablePrefix + defaultTableName;
	}
	db.SingularTable(true)

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

  	// defer db.CloseDB()
}

func CloseDB() {
	defer db.Close()
}