package models

type Article struct {
	Model
	TagId      *string `json:"tag_id"`
	Title      *string `json:"title" binding:required"`
	Desc       *string `json:"desc"`
	Content    *string `json:"content binding:required"`
	CreateBy   *string `json:"create_by"`
	ModifiedBy *string `json:"modified_by"`
	State      *int    `json:"state"`
}

func GetArticle(id int) (Articles []Article) {
	db.Where("id = ?", id).First(&Articles)
	return
}

func GetArticles() (Articles []Article) {
	db.Find(&Articles)
	return
}
