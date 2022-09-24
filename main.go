package main

import (
	"go_md_blog/cron"
	"go_md_blog/router"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Llongfile)

	cron.Init()
	root := router.Init()
	root.Run(":80")
}
