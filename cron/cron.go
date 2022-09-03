package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// Init ...
func Init() {
	syncIndex()

	job := cron.New()
	job.AddFunc("* * * * *", syncMd)
	job.AddFunc("* * * * *", syncIndex)
	job.AddFunc("* * * * *", func() { fmt.Println("test ...") })
	job.Start()
}
