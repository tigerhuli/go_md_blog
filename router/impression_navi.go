package router

import (
	"go_md_blog/cache"
	"go_md_blog/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

// impressionNavi 拉取文章导航
func impressionNavi(c *gin.Context) {
	gin_h := gin.H{}
	gin_h["impressions"] = cache.Impressions[0:constant.ImpressionNaviMoreSize]
	c.HTML(http.StatusOK, "impression_navi.html", gin_h)
}
