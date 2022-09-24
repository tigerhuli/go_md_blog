package cache

import (
	"go_md_blog/constant"
	"os"
	"sync/atomic"

	"gopkg.in/yaml.v2"
)

// PageView 页面访问量
type PageView struct {
	Total int64 `yaml:"total"` // 总访问量
}

// pageViewValue 页面访问缓存
var PageViewTotal atomic.Int64

func init() {
	content, err := os.ReadFile(constant.PageViewPath)
	if err != nil {
		panic(err.Error())
	}

	var page_view PageView
	err = yaml.Unmarshal(content, &page_view)
	if err != nil {
		panic(err.Error())
	}

	PageViewTotal.Store(page_view.Total)
}
