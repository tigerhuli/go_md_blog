package cron

import (
	"fmt"
	"go_md_blog/cache"
	"go_md_blog/constant"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// syncArticleTree 同步首页数据
func syncArticleTree() {
	html_dir := constant.ArticlesHtmlPath
	content, err := genIndexDfs(html_dir, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cache.ArticleTree = content
}

// genIndexDfs ...
func genIndexDfs(dir string, layer int) (string, error) {
	items, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("read %s failed", dir)
	}

	var str_builder strings.Builder
	for _, item := range items {
		if item.IsDir() {
			dir_content := genDirContent(layer, fmt.Sprintf("%s/%s", dir, item.Name()))
			str_builder.WriteString(dir_content)

			content, err := genIndexDfs(fmt.Sprintf("%s/%s", dir, item.Name()), layer+1)
			if err != nil {
				log.Printf("getMdPaths failed: %s", err)
				continue
			}
			str_builder.WriteString(content)
			continue
		}

		layer_content := genArticleContent(layer, fmt.Sprintf("%s/%s", dir, item.Name()))
		str_builder.WriteString(layer_content)
	}

	return str_builder.String(), nil
}

// genArticleContent 生成层级信息
func genArticleContent(layer int, path string) string {
	url := strings.TrimPrefix(path, constant.ArticlesHtmlPath)
	url = fmt.Sprintf("http://tigerhuli.com/article/%s", url)
	content := `<h${layer}><a href="${url}" target="_blank">${name}</a></h${layer}>`
	content = strings.ReplaceAll(content, "${url}", url)

	name := filepath.Base(path)
	name = strings.Split(name, ".")[0]
	content = strings.ReplaceAll(content, "${name}", name)

	content = strings.ReplaceAll(content, "${layer}", fmt.Sprintf("%d", layer))

	return content
}

// genDirContent 生成层级信息
func genDirContent(layer int, path string) string {
	content := `<h${layer}>${name}</a></h${layer}>`
	name := filepath.Base(path)
	name = strings.Split(name, ".")[0]
	content = strings.ReplaceAll(content, "${name}", name)

	content = strings.ReplaceAll(content, "${layer}", fmt.Sprintf("%d", layer))

	return content
}
