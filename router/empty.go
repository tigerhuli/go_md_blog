package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// empty 空请求返回结果
func empty(c *gin.Context) {
	c.HTML(http.StatusOK, "empty.html", gin.H{})
}
