package router

import (
	"go_md_blog/cache"
	"net/http"

	"github.com/gin-gonic/gin"
)

// impressionNavi 拉取文章导航
func impressionNavi(c *gin.Context) {
	gin_h := gin.H{}
	gin_h["impressions"] = cache.Impressions
	c.HTML(http.StatusOK, "impression_navi.html", gin_h)
}
