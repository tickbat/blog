package models

import "time"

type Article struct {
	Model
	TagId      int       `json:"tag_id"`
	Tag        *Tag      `json:"tag" binding:"-"`
	Title      string    `json:"title" binding:"required"`
	Desc       string    `json:"desc"`
	Content    string    `json:"content" binding:"required"`
	CreateDby  string    `json:"created_by"`
	ModifieDby string    `json:"modified_by"`
	State      int       `json:"state" binding:"required,eq=1|eq=2"`
	ImageUrl   string    `json:"image_url" binding:"required,max=255"`
	DeletedAt  time.Time `json:"-"`
}

type QueryArticle struct {
	TagId int    `json:"tag_id"`
	Title string `json:"title"`
	State int    `json:"state"`
}
