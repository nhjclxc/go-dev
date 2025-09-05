package main

import (
	"context"
	"framework_09_cron/cron"
	"framework_09_cron/cron/task"
	"time"
)

func main() {
	cronTaskManager := cron.NewCronTaskManager()
	cronTaskManager.AddTask(task.NewTask01(5))
	cronTaskManager.Start(context.Background())

	time.Sleep(1 * time.Minute)
}
