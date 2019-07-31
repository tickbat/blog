package models

import (
	"time"
)

type Article struct {
	Model
	TagId      *int       `json:"tag_id"`
	Tag        *Tag       `json:"tag" binding:"-"`
	Title      *string    `json:"title" binding:"required"`
	Desc       *string    `json:"desc"`
	Content    *string    `json:"content" binding:"required"`
	CreateDby  *string    `json:"created_by"`
	ModifieDby *string    `json:"modified_by"`
	State      *int       `json:"state" binding:"required,eq=1|eq=2"`
	ImageUrl   *string    `json:"image_url" binding:"required,max=255"`
	DeletedAt  *time.Time `json:"-"`
}

type QueryArticle struct {
	TagId *string `json:"tag_id"`
	Title *string `json:"title"`
	State *int    `json:"state"`
}

func GetArticle(id int) (article Article) {
	Db.Where("id = ?", id).First(&article)
	Db.Model(&article).Related(&article.Tag)
	return
}

func GetArticles(page, size int, conditions QueryArticle) (Articles []Article) {
	Db.Preload("Tag").Where(&conditions).Offset(page).Limit(size).Find(&Articles)
	return
}

func GetArticlesTotal(conditions QueryArticle) (count int) {
	Db.Model(&Article{}).Where(&conditions).Count(&count)
	return
}

func ExistArticleByID(id int) bool {
	var article Article
	Db.Select("id").Where("id = ?", id).First(&article)
	if article.ID == 0 {
		return false
	}
	return true
}

func AddArticle(article Article) {
	Db.Create(&article)
}

func EditArticle(article Article) {
	Db.Model(&article).Update(article)
}

func DeleteArticle(id int) {
	var article Article
	article.ID = id
	Db.Delete(&article)
}

func ClearAllArticle() error {
	if err := Db.Unscoped().Where("deleted_at is not null").Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
