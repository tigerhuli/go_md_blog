package cron

import (
	"go_md_blog/cache"
	"go_md_blog/constant"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// flushPageView 保存页面访问结果
func flushPageView() {
	var page_view cache.PageView
	page_view.Total = cache.PageViewTotal.Load()

	log.Printf("flush page view: %+v", page_view)

	content, err := yaml.Marshal(page_view)
	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile(constant.PageViewPath, content, 0666)
	if err != nil {
		panic(err.Error())
	}
}
