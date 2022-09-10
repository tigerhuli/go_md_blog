package router

import (
	"go_md_blog/cache"
	"net/http"

	"github.com/gin-gonic/gin"
)

// articleNavi 拉取文章导航
func articleNavi(c *gin.Context) {
	gin_h := gin.H{}
	gin_h["article_mm"] = cache.ArticleMm
	gin_h["article_tree"] = cache.ArticleTree
	c.HTML(http.StatusOK, "article_navi.html", gin_h)
}
