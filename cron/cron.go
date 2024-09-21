package cron

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"notifyy.app/backend/helpers"
)

func StartCron() {
	log.Info("Creating new cron scheduler...")
	c := cron.New()

	c.AddFunc("*/30 * * * *", func() { helpers.HourlyCron() })
	c.Start()
	select {}
}
