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
	s "trains/service"
)

func main() {
	//go run . -feature stations.map waterloo euston 5
	//go run . -feature jungle-desert.map jungle desert 5
	//go run . -feature smallAndLarge.map small large 9
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	conf, err := c.FlagHandling()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	appData, errs := c.DataConfiguration(conf)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
		os.Exit(1)
	}

	//wg for 1 goroutine that will run the algorithm
	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan struct{})

	//running algorithm in separate goroutine
	go func() {
		defer wg.Done()
		defer close(done)

		s.AlgorithmRunner(ctx, appData)
	}()

	select {
	//if ctrl + c was pressed
	case <-ctx.Done():
		log.Println("Shutdown requested.")
		//means that program wont exit until every goroutine that called Add(1) ti wg has called Done()
		wg.Wait()

	//when done is closed - exiting
	case <-done:
		log.Println("Program ended the work itself.")
	}

	log.Println("Exiting the program.")
}
