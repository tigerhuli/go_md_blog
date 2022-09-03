package router

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

// article 拉取文章
func article(c *gin.Context) {
	article_path := c.Request.URL.Path
	article_path = strings.TrimPrefix(article_path, "/article/")
	article_path = strings.TrimSuffix(article_path, "/")

	fmt.Println("get article path", article_path)
	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./output/%s", article_path))
	if err != nil {
		panic(err)
	}

	c.Writer.Write(fileBytes)
}
