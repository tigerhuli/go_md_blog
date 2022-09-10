package router

import (
	"fmt"
	"go_md_blog/constant"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// impression 拉取文章
func impression(c *gin.Context) {
	impression_path := c.Request.URL.Path
	impression_path = strings.TrimPrefix(impression_path, "/impression/")
	impression_path = strings.TrimSuffix(impression_path, "/")

	fmt.Println("get impression path", impression_path)
	impression_content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", constant.ImpressionsHtmlPath, impression_path))
	if err != nil {
		panic(err)
	}

	gin_h := gin.H{}
	gin_h["impression_content"] = string(impression_content)
	c.HTML(http.StatusOK, "impression.html", gin_h)
}
