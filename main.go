package main

import (
	"go_md_blog/cron"
	"go_md_blog/router"
)

func main() {
	cron.Init()
	root := router.Init()
	root.Run(":80")
}
