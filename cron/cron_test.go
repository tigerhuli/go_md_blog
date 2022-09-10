package cron

import "testing"

func TestSyncMd(t *testing.T) {
	syncMdToHtml()
}

func TestSyncIndex(t *testing.T) {
	syncArticleTree()
}

func TestSyncMmIndex(t *testing.T) {
	syncArticleMm()
}
