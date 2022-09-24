package cron

import (
	"testing"
	"time"
)

func TestSyncMd(t *testing.T) {
	syncArticles()
}

func TestSyncIndex(t *testing.T) {
	syncArticleTree()
}

func TestSyncMmIndex(t *testing.T) {
	syncArticleMm()
}

func TestClearExpiredHtml(t *testing.T) {
	start_time := time.Now()
	clearExpiredHtmlWithDir(start_time, "../content_html/articles")
}
