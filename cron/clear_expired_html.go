package cron

import (
	"go_md_blog/constant"
	"log"
	"os"
	"path"
	"time"
)

// clearExpiredHtml 清理过期的HTML文件
func clearExpiredHtml(start_time time.Time) {
	dirs := []string{
		constant.ArticlesHtmlPath,
		constant.ImpressionsHtmlPath,
	}

	for _, dir := range dirs {
		clearExpiredHtmlWithDir(start_time, dir)
	}
}

// clearExpiredHtmlWithDir 清理指定路径过期的HTML文件
func clearExpiredHtmlWithDir(start_time time.Time, input_dir string) {
	items, err := os.ReadDir(input_dir)
	if err != nil {
		log.Printf("read %s failed", input_dir)
		return
	}

	for _, item := range items {
		if item.IsDir() {
			dir := path.Join(input_dir, item.Name())
			file_stat, err := os.Stat(dir)
			if err != nil {
				log.Printf("get file %s stat failed: %s", dir, err.Error())
				continue
			}

			if file_stat.ModTime().After(start_time) {
				clearExpiredHtmlWithDir(start_time, dir)
			} else {
				err := os.RemoveAll(dir)
				if err != nil {
					log.Printf("remove dir %s failed: %s", dir, err.Error())
					return
				}
			}
			continue
		}

		file_path := path.Join(input_dir, item.Name())
		file_stat, err := os.Stat(file_path)
		if err != nil {
			log.Printf("get file stat filed: %s", err.Error())
			return
		}

		if file_stat.ModTime().After(start_time) {
			continue
		}

		err = os.Remove(file_path)
		if err != nil {
			log.Printf("remove file failed: %s", err.Error())
			return
		}
	}
}
