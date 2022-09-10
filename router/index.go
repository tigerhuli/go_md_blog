package router

import (
	"go_md_blog/cache"
	"net/http"

	"github.com/gin-gonic/gin"
)

// index 空请求返回结果
func index(c *gin.Context) {
	gin_h := gin.H{}
	gin_h["articles"] = cache.Articles
	c.HTML(http.StatusOK, "index.html", gin_h)
}
