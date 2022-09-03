package router

import "github.com/gin-gonic/gin"

// Init 初始化gin接口
func Init() *gin.Engine {
	root := gin.Default()
	root.GET("/image/:name", image)
	// root.Static("/image", "./input")
	return root
}
