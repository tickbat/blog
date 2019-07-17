package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       *string `json:"name" binding:"required"`
	CreatedBy  *string `json:"create_by"`
	ModifiedBy *string `json:"modified_by"`
	State      *int    `json:"state" binding:"required,eq=1|eq=2"`
}

type QueryTag struct {
	Name  *string `json:"name"`
	State *int    `json:"state" binding:"eq=1|eq=2"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagsTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID == nil {
		return false
	}

	return true
}

func AddTag(tag Tag) bool {
	db.Create(&tag)
	return true
}

func EditTag(tag Tag) {
	db.Model(&tag).Update(tag)
}

func DeleteTag(id int) {
	tag := new(Tag)
	tag.ID = &id
	db.Delete(tag)
}
