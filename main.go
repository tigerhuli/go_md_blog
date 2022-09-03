package main

import "go_md_blog/router"

func main() {
	root := router.Init()
	root.Run(":80")
}
