package router

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// Init 初始化gin接口
func Init() *gin.Engine {
	root := gin.Default()
	root.SetFuncMap(template.FuncMap{ // 取消嵌入html后的双引号
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	root.LoadHTMLGlob("./templates/*")

	root.GET("/", mmIndex)
	root.GET("/mm_index", mmIndex)
	root.GET("/index", index)
	root.GET("/image/*action", image)
	root.GET("/article/*action", article)
	root.StaticFile("/favicon_zzz", "./resources/favicon.png")

	return root
}
