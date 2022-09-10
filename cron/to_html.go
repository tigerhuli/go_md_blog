package cron

import (
	"bytes"
	"fmt"
	"go_md_blog/constant"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"

	mermaid "github.com/abhinav/goldmark-mermaid"
)

// toHTML 将markdown转换为html
func toHTML(md_path, html_path string, md_content []byte) {
	dir_path := filepath.Dir(html_path)
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		err := os.MkdirAll(dir_path, 0666)
		if err != nil {
			log.Println(err.Error())
		}
	}

	var buf bytes.Buffer
	if err := goldmark.New(goldmark.WithExtensions(&mermaid.Extender{})).Convert(md_content, &buf); err != nil {
		log.Println(err.Error())
		return
	}

	html_content := buf.Bytes()
	html_content = replaceHtmlImage(md_path, html_content)
	err := ioutil.WriteFile(html_path, html_content, 0666)
	if err != nil {
		log.Println(err.Error())
	}
}

// replaceHtmlImage 替换html中的image
func replaceHtmlImage(md_path string, content []byte) []byte {
	image_reg := regexp.MustCompile(`<img[\s\S]+?>`)
	matches := image_reg.FindAll(content, -1)

	tag_map := make(map[string]string)
	for _, match := range matches {
		tag := string(match)
		if _, ok := tag_map[tag]; ok {
			continue
		}

		path_reg := regexp.MustCompile(`(src=")(.*?)(")`)
		sub_matches := path_reg.FindSubmatch(match)
		if len(sub_matches) < 4 {
			continue
		}

		path := string(sub_matches[2])
		if strings.Contains(path, "http") { // ignore http link
			continue
		}

		dir := filepath.Dir(md_path)
		new_path := filepath.Join(dir, path) // 处理相对路径
		new_path = strings.TrimPrefix(new_path, constant.ContentPath)
		new_path = fmt.Sprintf(`%s/image/%s`, constant.BaseUrl, new_path)
		new_tag := strings.Replace(tag, path, new_path, 1)

		tag_map[tag] = new_tag
	}

	for tag, new_tag := range tag_map {
		content = bytes.ReplaceAll(content, []byte(tag), []byte(new_tag))
	}

	return content
}
