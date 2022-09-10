package cron

import "testing"

func TestSyncMd(t *testing.T) {
	syncArticle()
}

func TestSyncIndex(t *testing.T) {
	syncArticleTree()
}

func TestSyncMmIndex(t *testing.T) {
	syncArticleMm()
}

func TestRegExp(t *testing.T) {

}
