package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	c "trains/cli"
	p "trains/pathfinder"
	s "trains/scheduler"
)

func main() {
	// go run . -feature file start end 7
	conf, err := c.FlagHandling()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loaded trains:", conf.TrainNumb)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	res := p.PathfinderAlgorithm(ctx)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.SchedulerAlgorithm(ctx, res)
	}()

	<-ctx.Done()

	log.Println("Shutdown requested")

	wg.Wait()
}
