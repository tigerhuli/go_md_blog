package tohtml

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"

	mermaid "github.com/abhinav/goldmark-mermaid"
)

// ToHTML 将markdown转换为html
func ToHTML(md_path, html_path string) {
	md_content, err := ioutil.ReadFile(md_path)
	if err != nil {
		fmt.Println(err.Error())
	}

	var buf bytes.Buffer
	if err := goldmark.New(goldmark.WithExtensions(&mermaid.Extender{})).Convert(md_content, &buf); err != nil {
		panic(err)
	}

	html_content := buf.Bytes()
	html_content = replaceHtmlImage(html_content)
	ioutil.WriteFile(html_path, html_content, 0666)
}

// replaceHtmlImage 替换html中的image
func replaceHtmlImage(content []byte) []byte {
	image_reg := regexp.MustCompile(`<img[\s\S]+?>`)
	matches := image_reg.FindAll(content, -1)
	fmt.Println(len(matches))

	tag_map := make(map[string]string)
	for _, match := range matches {
		tag := string(match)
		if _, ok := tag_map[tag]; ok {
			continue
		}
		fmt.Println(tag)
		path_reg := regexp.MustCompile(`(src=")(.*?)(")`)
		sub_matches := path_reg.FindSubmatch(match)
		if len(sub_matches) < 4 {
			continue
		}
		path := string(sub_matches[2])
		fmt.Println(path)

		new_path := fmt.Sprintf(`http://tigerhuli.com/image/%s`, path)
		new_tag := strings.Replace(tag, path, new_path, 1)

		tag_map[tag] = new_tag
	}

	for tag, new_tag := range tag_map {
		content = bytes.ReplaceAll(content, []byte(tag), []byte(new_tag))
	}

	fmt.Println(string(content))
	return content
}
