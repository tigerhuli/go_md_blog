package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// image 获取图片
func image(c *gin.Context) {
	image_name := c.Param("name")
	fmt.Println("get image name", image_name)
}
