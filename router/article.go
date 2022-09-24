package router

import (
	"fmt"
	"go_md_blog/constant"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// article 拉取文章
func article(c *gin.Context) {
	article_path := c.Request.URL.Path
	article_path = strings.TrimPrefix(article_path, "/article/")
	article_path = strings.TrimSuffix(article_path, "/")

	fmt.Println("get article path", article_path)
	article_content, err := os.ReadFile(fmt.Sprintf("%s/%s", constant.ArticlesHtmlPath, article_path))
	if err != nil {
		panic(err)
	}

	gin_h := gin.H{}
	gin_h["article_content"] = string(article_content)
	c.HTML(http.StatusOK, "article.html", gin_h)
}
