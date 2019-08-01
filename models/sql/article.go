package sql

import (
	"blog/models"
	"github.com/jinzhu/gorm"
)

func GetArticle(id int) (models.Article, error) {
	var article models.Article
	if err := models.Db.Where("id = ?", id).First(&article).Error; err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	if err := models.Db.Model(&article).Related(&article.Tag).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetArticles(conditions *models.QueryArticle, page, size int) ([]models.Article, error) {
	var articles []models.Article
	if err := models.Db.Preload("Tag").Where(conditions).Offset(page).Limit(size).Find(&articles).Error; err != nil && err != gorm.ErrRecordNotFound {
		return articles, err
	}
	return articles, nil
}

func GetArticlesTotal(conditions models.QueryArticle) (int, error) {
	var count int
	err := models.Db.Model(&models.Article{}).Where(&conditions).Count(&count).Error
	return count, err
}

func ExistArticleByID(id int) (bool, error) {
	var article models.Article
	err := models.Db.Select("id").Where("id = ?", id).First(&article).Error
	if article.ID == 0 {
		return false, err
	}
	return true, err
}

func AddArticle(article models.Article) error {
	err := models.Db.Create(&article).Error
	return err
}

func EditArticle(article models.Article) error {
	err := models.Db.Model(&article).Update(article).Error
	return err
}

func DeleteArticle(id int) error {
	var article models.Article
	article.ID = id
	err := models.Db.Delete(&article).Error
	return err
}

func ClearAllArticle() error {
	err := models.Db.Unscoped().Where("deleted_at is not null").Delete(&models.Article{}).Error
	return err
}
