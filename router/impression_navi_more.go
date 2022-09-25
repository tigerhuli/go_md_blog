package router

import (
	"go_md_blog/cache"
	"go_md_blog/constant"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// impressionNaviMore 展示更多impression内容
func impressionNaviMore(c *gin.Context) {
	var page int
	page_str, ok := c.GetQuery("page")
	if ok {
		tmp_page, err := strconv.Atoi(page_str)
		if err != nil {
			log.Printf("convert page failed: %s", err.Error())
			page = 0
		}
		page = tmp_page
	} else {
		page = 0
	}

	batch_size := constant.ImpressionNaviMoreSize
	start := page * batch_size
	if start >= len(cache.Impressions) {
		c.HTML(http.StatusNotFound, "no_more.html", nil)
		return
	}

	end := start + batch_size
	if end > len(cache.Impressions) {
		end = len(cache.Impressions)
	}

	log.Printf("get page: %d, start: %d, end: %d", page, start, end)
	gin_h := gin.H{}
	gin_h["impressions"] = cache.Impressions[start:end]
	c.HTML(http.StatusOK, "impression_navi_more.html", gin_h)
}
