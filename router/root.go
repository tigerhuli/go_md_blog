package router

import "github.com/gin-gonic/gin"

// Init 初始化gin接口
func Init() *gin.Engine {
	root := gin.Default()
	root.GET("/image/*action", image)
	root.GET("/article/*action", article)
	root.StaticFile("favicon.ico", "./resources/favicon.png")

	return root
}
