package main

import (
	"bp14_task_queue/queue"
	"bp14_task_queue/queue/priority"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

var delayQueue *priority.PriorityQueue

func init() {
	ctx := context.Background()
	delayQueue = priority.NewPriorityQueue(ctx, 3)
}

// 任务生产者
func main() {
	ctx := context.Background()

	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i)
	}

	port := os.Getenv("PORT")
	producerName := os.Getenv("NAME")
	if port == "" {
		port = "8080"
	}

	e := gin.Default()

	// http://127.0.0.1:8080/api/sendTask?id=1&name=task1&delay=10&priorityTmp=1
	// http://127.0.0.1:8080/api/sendTask?id=2&name=task21&delay=10&priorityTmp=2
	// http://127.0.0.1:8080/api/sendTask?id=3&name=task22&delay=10&priorityTmp=2
	e.GET("/api/sendTask", func(c *gin.Context) {
		var req queue.QueueTask
		err := c.ShouldBindQuery(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Delay = time.Duration(req.DelayTmp) * time.Second
		req.JoinTime = time.Now()

		fmt.Println("req:", req)

		err = delayQueue.Enqueue(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("\n\n\n")
		c.JSON(200, gin.H{
			"message": "success: " + req.Name + time.Now().String(),
		})

	})
	log.Printf("🚀 %s on http://0.0.0.0:%s", producerName, port)
	e.Run(":" + port)

}

func worker(ctx context.Context, id int) {

	err := delayQueue.AddWorker()
	if err != nil {
		fmt.Printf("Worker %d <%s>\n", id, err)
		return
	}
	defer delayQueue.DoneWorker()

	fmt.Printf("Worker %d 启动\n", id)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d 接收到退出信号，停止工作\n", id)
			return
		default:
			dequeue, err := delayQueue.Dequeue(ctx)
			if err != nil {
				continue
			}

			// 模拟任务执行失败
			if id == dequeue.Id {
				reson := fmt.Errorf("Worker %d - %d <workerId==dequeue.Id 不允许执行>\n", id, dequeue.Id)
				log.Println(reson)
				delayQueue.Nack(ctx, dequeue, 2*time.Second, reson)
				// 跳过任务的执行
				continue
			}

			log.Printf("Worker %d 正在工作: %#v \n", id, dequeue)
			time.Sleep(3 * time.Second) // 模拟任务执行
		}
	}
}
