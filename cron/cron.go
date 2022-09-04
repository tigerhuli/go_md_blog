package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// Init ...
func Init() {
	syncIndex()
	syncMmIndex()

	job := cron.New()
	cron_daily := func() {
		syncMd()
		syncIndex()
		syncMmIndex()
	}
	job.AddFunc("0 0 * * *", cron_daily)
	job.AddFunc("* * * * *", func() { fmt.Println("test ...") })
	job.Start()
}
