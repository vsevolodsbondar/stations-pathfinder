package scheduler

import (
	"context"
	"log"
	"time"
)

func SchedulerAlgorithm(ctx context.Context, paths string) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Scheduler is stopping the work...")
			return

		default:
			log.Println("Scheduling...")
			time.Sleep(5000 * time.Millisecond)
		}
	}
}
