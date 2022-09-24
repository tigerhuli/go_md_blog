package cron

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// getFilePaths 拉取某种类型文件的路径
func getFilePaths(input_dir string, file_formats []string) ([]string, error) {
	items, err := os.ReadDir(input_dir)
	if err != nil {
		log.Printf("read %s failed", input_dir)
		return nil, err
	}

	var file_paths []string
	for _, item := range items {
		if item.IsDir() {
			inner_md_paths, err := getMdPaths(fmt.Sprintf("%s/%s", input_dir, item.Name()))
			if err != nil {
				log.Printf("getMdPaths failed: %s", err)
				continue
			}
			file_paths = append(file_paths, inner_md_paths...)
			continue
		}

		for _, file_format := range file_formats {
			if strings.HasSuffix(item.Name(), file_format) {
				file_paths = append(file_paths, fmt.Sprintf("%s/%s", input_dir, item.Name()))
			}
		}
	}

	return file_paths, nil
}
