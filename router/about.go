package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// about 空请求返回结果
func about(c *gin.Context) {
	gin_h := gin.H{}
	c.HTML(http.StatusOK, "about.html", gin_h)
}
