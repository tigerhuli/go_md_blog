package router

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// image 获取图片
func image(c *gin.Context) {
	image_name := c.Param("name")

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./input/%s", image_name))
	if err != nil {
		panic(err)
	}

	c.Writer.Write(fileBytes)
}
