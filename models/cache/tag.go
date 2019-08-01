package cache

import (
	"blog/models"
	"strconv"
	"strings"
)

func GetTagsKey(t *models.QueryTag, pageNum, pageSize int) string {
	keys := []string{
		"TAG-LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State != 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if pageNum != 0 {
		keys = append(keys, strconv.Itoa(pageNum))
	}
	if pageSize != 0 {
		keys = append(keys, strconv.Itoa(pageSize))
	}

	return strings.Join(keys, "_")
}
