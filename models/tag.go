package models

type Tag struct {
	Model
	Name       string `json:"name" binding:"required"`
	CreateDby  string `json:"create_by"`
	ModifieDby string `json:"modified_by"`
	State      int    `json:"state" binding:"required,eq=1|eq=2"`
}

type QueryTag struct {
	Name     string `form:"name"`
	State    int    `form:"state" binding:"omitempty,eq=1|eq=2"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"PageSize"`
}

func (q QueryTag) TableName() string {
	return "tag"
}
