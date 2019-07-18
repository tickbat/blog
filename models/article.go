package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagId      *int    `json:"tag_id"`
	Tag        Tag     `json:"tag" binding:"-"`
	Title      *string `json:"title" binding:"required"`
	Desc       *string `json:"desc"`
	Content    *string `json:"content" binding:"required"`
	CreatedBy  *string `json:"created_by"`
	ModifiedBy *string `json:"modified_by"`
	State      *int    `json:"state" binding:"required"`
}

type QueryArticle struct {
	TagId *string `json:"tag_id"`
	Title *string `json:"title"`
	State *int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func GetArticles(page, size int, conditions QueryArticle) (Articles []Article) {
	db.Preload("Tag").Where(&conditions).Offset(page).Limit(size).Find(&Articles)
	return
}

func GetArticlesTotal(conditions QueryArticle) (count int) {
	db.Model(&Article{}).Where(&conditions).Count(&count)
	return
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID == nil {
		return false
	}
	return true
}

func AddArticle(article Article) {
	db.Create(&article)
}

func EditArticle(article Article) {
	db.Model(&article).Update(article)
}

func DeleteArticle(id int) {
	var article Article
	article.ID = &id
	db.Delete(&article)
}
