package main

import (
	cronTask "best_practice_06_cron/cron_task"
	"log"
	"time"
)

func main() {

	ec := make(chan bool)

	manager := cronTask.NewCronTaskManager(true)
	manager.Start()
	defer manager.Stop()

	manager.AddTask("helloTask", "*/10 * * * * *", func() {
		log.Println("Hello, Cron!")
	})

	time.Sleep(35 * time.Second)

	tasks := manager.ListTasks()
	log.Println("Tasks:", tasks)

	<-ec

}
