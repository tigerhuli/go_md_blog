package cron

import (
	"bytes"
	"fmt"
	"go_md_blog/cache"
	"go_md_blog/constant"
	"go_md_blog/model"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// syncArticles 同步article到html
func syncArticles() {
	input_dir := constant.ArticlesPath
	md_paths, err := getMdPaths(input_dir)
	if err != nil {
		fmt.Println(err.Error())
	}

	var articles []model.Article
	for _, md_path := range md_paths {
		html_path := strings.TrimPrefix(md_path, input_dir)
		html_path = strings.Split(html_path, ".")[0]
		html_path = fmt.Sprintf("%s%s.html", constant.ArticlesHtmlPath, html_path)

		md_content, err := os.ReadFile(md_path)
		if err != nil {
			log.Println(err.Error())
			return
		}

		article, md_content, err := getArticleFromMd(md_path, html_path, md_content)
		if err != nil {
			log.Printf("getArticleFromMdPath failed: %s", err.Error())
			continue
		}

		articles = append(articles, article)

		toHTML(md_path, html_path, md_content)
	}

	cache.Articles = articles
}

// getArticleFromMd 从md文件中提起article信息
func getArticleFromMd(md_path, html_path string, md_content []byte) (model.Article, []byte, error) {
	r := regexp.MustCompile(`(^---)([\s\S]*?)(---)`)
	submatches := r.FindSubmatch([]byte(md_content))
	if len(submatches) == 0 {
		return model.Article{}, nil, fmt.Errorf("no article info find")
	}

	md_content = bytes.Replace(md_content, submatches[0], nil, 1)

	var article model.Article
	err := yaml.Unmarshal(submatches[2], &article)
	if err != nil {
		return model.Article{}, nil, fmt.Errorf("unmarshal article info failed: %s", err)
	}

	if len(article.Title) == 0 {
		title := filepath.Base(md_path)
		title = strings.Split(title, ".")[0]
		article.Title = title
	}

	if len(article.Url) == 0 {
		url := strings.TrimPrefix(html_path, constant.ArticlesHtmlPath)
		url = fmt.Sprintf("http://tigerhuli.com/article/%s", url)
		article.Url = url
	}

	return article, md_content, nil
}

// getMdPaths 拉取markdown列表
func getMdPaths(input_dir string) ([]string, error) {
	items, err := os.ReadDir(input_dir)
	if err != nil {
		log.Printf("read %s failed", input_dir)
	}

	var md_paths []string
	for _, item := range items {
		if item.IsDir() {
			inner_md_paths, err := getMdPaths(fmt.Sprintf("%s/%s", input_dir, item.Name()))
			if err != nil {
				log.Printf("getMdPaths failed: %s", err)
				continue
			}
			md_paths = append(md_paths, inner_md_paths...)
			continue
		}

		if strings.HasSuffix(item.Name(), ".md") || strings.HasSuffix(item.Name(), ".markdown") {
			md_paths = append(md_paths, fmt.Sprintf("%s/%s", input_dir, item.Name()))
		}
	}

	return md_paths, nil
}
