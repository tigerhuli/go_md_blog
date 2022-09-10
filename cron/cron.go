package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// Init ...
func Init() {
	syncMdToHtml()
	syncArticleTree()
	syncArticleMm()

	job := cron.New()
	cron_daily := func() {
		syncMdToHtml()
		syncArticleTree()
		syncArticleMm()
	}
	job.AddFunc("0 0 * * *", cron_daily)
	job.AddFunc("* * * * *", func() { fmt.Println("test ...") })
	job.Start()
}
