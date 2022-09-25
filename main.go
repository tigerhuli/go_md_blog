package main

import (
	"go_md_blog/cron"
	"go_md_blog/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFlags(log.Ldate | log.Llongfile)

	cron.Init()
	root := router.Init()

	go func() {
		root.Run(":80")
	}()

	// graceful exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cron.FlushPageView()
}
