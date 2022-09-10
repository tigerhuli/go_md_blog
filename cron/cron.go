package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// Init ...
func Init() {
	syncArticles()
	syncArticleTree()
	syncArticleMm()
	syncImpressions()

	job := cron.New()
	cron_daily := func() {
		syncArticles()
		syncArticleTree()
		syncArticleMm()
		syncImpressions()
	}
	job.AddFunc("0 0 * * *", cron_daily)
	job.AddFunc("* * * * *", func() { fmt.Println("test ...") })
	job.Start()
}
