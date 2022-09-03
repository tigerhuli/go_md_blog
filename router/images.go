package router

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

// image 获取图片
func image(c *gin.Context) {
	image_path := c.Request.URL.Path
	image_path = strings.TrimPrefix(image_path, "/image/")
	image_path = strings.TrimSuffix(image_path, "/")

	fmt.Println("get image path", image_path)
	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./input/%s", image_path))
	if err != nil {
		panic(err)
	}

	c.Writer.Write(fileBytes)
}
