package cron

import (
	"notifyy.app/backend/helpers"
)

func StartCron() {
	helpers.HourlyCron()

}
