package service

import (
	"blog/models"
	"blog/models/cache"
	"blog/models/sql"
	"blog/pkg/gredis"
	"blog/pkg/logging"
	"encoding/json"
	"time"
)

func GetArticle(id int) (models.Article, error) {
	var articleList models.Article
	key := cache.GetArticleKey(id)
	val, err := gredis.Get(key)
	if err != nil && json.Unmarshal(val, &articleList) != nil {
		logging.Warn("get article from redis error:", err)
		data, err := sql.GetArticle(id)
		if err != nil {
			return data, err
		}
		if err := gredis.Set(key, data, time.Second*60); err != nil {
			logging.Warn("set article into redis error:", err)
		}
		return data, nil
	}
	return articleList, nil
}

func GetArticles(conditions *models.QueryArticle, pageNum, pageSize int) ([]models.Article, error) {
	var ArticleList []models.Article
	key := cache.GetArticlesKey(conditions, pageNum, pageSize)
	val, err := gredis.Get(key)
	if err != nil && json.Unmarshal(val, &ArticleList) != nil {
		logging.Warn("get tags from redis error:", err)
		data, err := sql.GetArticles(conditions, pageNum, pageSize)
		if err != nil {
			return data, err
		}
		if err := gredis.Set(key, data, time.Second*60); err != nil {
			logging.Warn("set tags into redis error:", err)
		}
		return data, nil
	}
	return ArticleList, nil
}

func GetArticlesTotal(conditions models.QueryArticle) (int, error) {
	return sql.GetArticlesTotal(conditions)
}

func ExistArticleByID(id int) (bool, error) {
	return sql.ExistArticleByID(id)
}

func AddArticle(article models.Article) error {
	return sql.AddArticle(article)
}

func EditArticle(article models.Article) error {
	return sql.EditArticle(article)
}

func DeleteArticle(id int) error {
	return sql.DeleteArticle(id)
}

func ClearAllArticle() error {
	return sql.ClearAllArticle()
}
