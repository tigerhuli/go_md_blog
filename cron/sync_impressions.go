package cron

import (
	"bytes"
	"fmt"
	"go_md_blog/cache"
	"go_md_blog/constant"
	"go_md_blog/model"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// syncImpressions 同步Impressions到html
func syncImpressions() {
	input_dir := constant.ImpressionsPath
	md_paths, err := getMdPaths(input_dir)
	if err != nil {
		fmt.Println(err.Error())
	}

	var impressions []model.Impression
	for _, md_path := range md_paths {
		html_path := strings.TrimPrefix(md_path, input_dir)
		html_path = strings.Split(html_path, ".")[0]
		html_path = fmt.Sprintf("%s%s.html", constant.ImpressionsHtmlPath, html_path)

		md_content, err := ioutil.ReadFile(md_path)
		if err != nil {
			log.Println(err.Error())
			return
		}

		impression, md_content, err := getImpressionFromMd(md_path, html_path, md_content)
		if err != nil {
			log.Printf("getArticleFromMdPath failed: %s", err.Error())
			continue
		}

		impressions = append(impressions, impression)

		toHTML(md_path, html_path, md_content)
	}

	cache.Impressions = impressions
}

// getArticleFromMd 从md文件中提起Impression信息
func getImpressionFromMd(md_path, html_path string, md_content []byte) (model.Impression, []byte, error) {
	r := regexp.MustCompile(`(^---)([\s\S]*?)(---)`)
	submatches := r.FindSubmatch([]byte(md_content))
	if len(submatches) == 0 {
		return model.Impression{}, nil, fmt.Errorf("no article info find")
	}

	md_content = bytes.Replace(md_content, submatches[0], nil, 1)

	var impression model.Impression
	err := yaml.Unmarshal(submatches[2], &impression)
	if err != nil {
		return model.Impression{}, nil, fmt.Errorf("unmarshal article info failed: %s", err)
	}

	if len(impression.Name) == 0 {
		name := filepath.Base(md_path)
		name = strings.Split(name, ".")[0]
		impression.Name = name
	}

	if len(impression.Url) == 0 {
		url := strings.TrimPrefix(html_path, constant.ArticlesHtmlPath)
		url = fmt.Sprintf("http://tigerhuli.com/article/%s", url)
		impression.Url = url
	}

	return impression, md_content, nil
}
