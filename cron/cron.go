package cron

import (
	"time"

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
		start_time := time.Now()

		syncArticles()
		syncArticleTree()
		syncArticleMm()
		syncImpressions()
		clearExpiredHtml(start_time)
	}
	job.AddFunc("0 0 * * *", cron_daily)
	job.AddFunc("* * * * *", FlushPageView)
	job.Start()
}
