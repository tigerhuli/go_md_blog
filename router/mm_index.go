package router

import (
	"go_md_blog/cache"
	"net/http"

	"github.com/gin-gonic/gin"
)

// mmIndex 空请求返回结果
func mmIndex(c *gin.Context) {
	gin_h := gin.H{}
	gin_h["mm_script"] = cache.IndexContent
	c.HTML(http.StatusOK, "index.html", gin_h)
}
