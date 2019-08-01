package service

import (
	"blog/models"
	"blog/models/cache"
	"blog/models/sql"
	"blog/pkg/file"
	"blog/pkg/gredis"
	"blog/pkg/logging"
	"blog/pkg/setting"
	"encoding/json"
	"strconv"
	"time"
	//"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

func GetTags(tag *models.QueryTag, pageNum, pageSize int) ([]models.Tag, error) {
	var tagList []models.Tag
	key := cache.GetTagsKey(tag, pageNum, pageSize)
	val, err := gredis.Get(key)
	if err != nil && json.Unmarshal(val, &tagList) != nil {
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
	return tagList, nil
	/*if err := json.Unmarshal(val, &tagList); err != nil {
		logging.Warn("get tags when parse json error:", err)
		data, err := sql.GetTags(tag, pageNum, pageSize)
		if err != nil {
			return data, err
		}
		if err := gredis.Set(key, data, time.Second * 60); err != nil {
			logging.Warn("set tags into redis error:", err)
		}
		return data, nil
	}*/
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

func ExistTagByID(id int) (bool, error) {
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

	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(v.CreatedOn),
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + ".xlsx"

	dirFullPath := setting.Excel.SavePath
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}

	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func ClearAllTag() error {
	if err := models.Db.Where("deleted_at != null").Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}
