package cache

import (
	"blog/models"
	"strconv"
	"strings"
)

func GetArticleKey(id int) string {
	return "ARTICLE" + "_" + strconv.Itoa(id)
}

func GetArticlesKey(article *models.QueryArticle, pageNum, pageSize int) string {
	keys := []string{
		"ARTICLE-LIST",
	}
	if article.State >= 0 {
		keys = append(keys, strconv.Itoa(article.State))
	}
	if article.Title != "" {
		keys = append(keys, article.Title)
	}
	if article.TagId > 0 {
		keys = append(keys, strconv.Itoa(article.TagId))
	}
	if pageNum > 0 {
		keys = append(keys, strconv.Itoa(pageNum))
	}
	if pageSize > 0 {
		keys = append(keys, strconv.Itoa(pageSize))
	}

	return strings.Join(keys, "_")
}
