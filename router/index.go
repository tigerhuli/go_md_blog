package router

import (
	"go_md_blog/cache"
	"go_md_blog/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

// index 空请求返回结果
func index(c *gin.Context) {
	gin_h := gin.H{}
	gin_h["articles"] = cache.Articles[0:constant.IndexMoreSize]
	c.HTML(http.StatusOK, "index.html", gin_h)
}
