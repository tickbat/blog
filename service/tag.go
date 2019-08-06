package service

import (
	"blog/models"
	"blog/models/cache"
	"blog/models/sql"
	"blog/pkg/gredis"
	"blog/pkg/logging"
	"blog/pkg/setting"
	"blog/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"strconv"
	"time"
)

func GetTags(tag *models.QueryTag, pageNum, pageSize int) ([]models.Tag, error) {
	var tagList []models.Tag
	key := cache.GetTagsKey(tag, pageNum, pageSize)
	val, err := gredis.Get(key)
	marshalErr := json.Unmarshal(val, &tagList)
	if err != nil && marshalErr != nil {
		var data []models.Tag
		logging.Warn("get tags from redis error:", err)
		if pageNum == 0 && pageSize == 0 {
			data, err = sql.GetTagsAll(tag)
		} else {
			data, err = sql.GetTags(tag, pageNum, pageSize)
		}
		if err != nil {
			return data, err
		}
		if err := gredis.Set(key, data, time.Second*60); err != nil {
			logging.Warn("set tags into redis error:", err)
		}
		return data, nil
	}
	fmt.Printf("%+v\n", tagList)
	return tagList, nil
}

func AddTag(tag models.Tag) error {
	return sql.AddTag(tag)
}

func EditTag(tag models.Tag) error {
	return sql.EditTag(tag)
}

func DeleteTag(id int) error {
	return sql.DeleteTag(id)
}

func ExistTagByID(id int) bool {
	return sql.ExistTagByID(id)
}

func GetTagsTotal(maps interface{}) (int, error) {
	return sql.GetTagsTotal(maps)
}

func ExportTag(tag models.QueryTag) (string, error) {
	tags, err := GetTags(&tag, 0, 0)
	if err != nil {
		return "", err
	}
	f := excelize.NewFile()
	f.NewSheet("标签信息")

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	for i, v := range titles {
		f.SetCellValue("sheet1", util.Axis(1, i), v)
	}
	for i, v := range tags {
		content := []interface{}{v.ID, v.Name, v.CreatedBy, v.CreatedOn, v.ModifiedBy, v.CreatedOn}
		for i2, v2 := range content {
			f.SetCellValue("sheet1", util.Axis(i+1, i2), v2)
		}
	}
	path := setting.Excel.SavePath
	times := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + times + ".xlsx"
	fullPath := path + filename
	if err := f.SaveAs(fullPath); err != nil {
		return "", err
	}
	return filename, nil
}

func ImportTag(r io.Reader) (error, int) {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err, 0
	}

	rows := xlsx.GetRows("标签信息")
	var tag models.Tag
	failNum := 0
	for index, row := range rows {
		if index > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			tag.Name = data[0]
			tag.State = 1
			tag.CreatedBy = data[2]
			if err := AddTag(tag); err != nil {
				logging.Error("import tag error", err)
				failNum++
			}
		}
	}
	return nil, failNum
}

func ClearAllTag() error {
	if err := models.Db.Where("deleted_at != null").Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}
