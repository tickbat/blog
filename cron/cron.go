package cron

import (
	"blog/models"
	"blog/pkg/logging"
	"github.com/robfig/cron"
)

const (
	clearTime = "0 0 12 * * *"
)

func init() {
	logging.Info("start cron")
	c := cron.New()
	c.AddFunc(clearTime, func() {
		if err := models.ClearAllArticle(); err != nil {
			logging.Error("cron ClearAllArticle error: " + err.Error())
		} else {
			logging.Info("cron ClearAllArticle success")
		}
	})
	c.Start()
}
