package cache_service

import (
	"blog/models"
	"strconv"
	"strings"
)

func GetTagsKey(t *models.QueryTag) string {
	keys := []string{
		"TAG-LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State != 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum != 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize != 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")
}
