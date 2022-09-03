package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// article 拉取文章
func article(c *gin.Context) {
	article_path := c.Request.URL.Path
	article_path = strings.TrimPrefix(article_path, "/article/")
	article_path = strings.TrimSuffix(article_path, "/")

	fmt.Println("get article path", article_path)
	md_content, err := ioutil.ReadFile(fmt.Sprintf("./output/%s", article_path))
	if err != nil {
		panic(err)
	}

	gin_h := gin.H{}
	gin_h["md_content"] = string(md_content)
	c.HTML(http.StatusOK, "article.html", gin_h)
}
