package cron

import (
	"fmt"
	"go_md_blog/tohtml"
	"io/ioutil"
	"strings"

	"log"
)

func syncMd() {
	input_dir := "../input"
	md_paths, err := getMdPaths(input_dir)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(md_paths)

	for _, md_path := range md_paths {
		html_path := strings.TrimPrefix(md_path, input_dir)
		html_path = strings.Split(html_path, ".")[0]
		html_path = fmt.Sprintf("../output%s.html", html_path)
		fmt.Println(html_path)
		tohtml.ToHTML(md_path, html_path)
	}
}

// getMds 拉取markdown列表
func getMdPaths(input_dir string) ([]string, error) {
	items, err := ioutil.ReadDir(input_dir)
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
