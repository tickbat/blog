package models

type Tag struct {
	Model
	Name       *string `json:"name" binding:"required"`
	CreatedBy  *string `json:"create_by"`
	ModifiedBy *string `json:"modified_by"`
	State      *int    `json:"state" binding:"required,eq=1|eq=2"`
}

type QueryTag struct {
	Name  *string `form:"name"`
	State *int    `form:"state" binding:"omitempty,eq=1|eq=2"`
}

func (q QueryTag) TableName() string {

	return "blog_tag"
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

func ClearAllTag() error {
	if err := db.Where("deleted_at != null").Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}