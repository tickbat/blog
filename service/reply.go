package service

import (
	"blog/models"
	"blog/models/sql"
)

func GetApplies(replay models.QueryReply) ([]models.Reply, error) {
	return sql.GetReplies(replay)
}

func AddReply(replay models.Reply) error {
	return sql.AddReply(replay)
}

func DeleteReply(id int) error {
	return sql.DeleteReply(id)
}

func ExistReplyByID(id int) bool {
	return sql.ExistReplyByID(id)
}
