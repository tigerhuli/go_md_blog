package router

import (
	"go_md_blog/cache"
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

// Init 初始化gin接口
func Init() *gin.Engine {
	root := gin.Default()
	root.SetTrustedProxies(nil)
	root.SetFuncMap(template.FuncMap{ // 取消嵌入html后的双引号
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	root.Use(addPageViewTotal())

	root.LoadHTMLGlob("./layout/*")

	root.GET("/", index)
	root.GET("/image/*action", image)
	root.Static("/static", "./static/")

	root.GET("/article_navi", articleNavi)
	root.GET("/article/*action", article)

	root.GET("/impression_navi", impressionNavi)
	root.GET("/impression/*action", impression)

	root.GET("/about", about)

	return root
}

// addPageViewTotal 自增页面访问
func addPageViewTotal() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/image") {
			return
		}

		if strings.HasPrefix(c.Request.URL.Path, "/static") {
			return
		}

		cache.PageViewTotal.Add(1)
	}
}
