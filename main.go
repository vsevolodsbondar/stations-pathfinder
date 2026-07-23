package main

import (
	"context"
	"fmt"
<<<<<<< HEAD
	"os"
	c "trains/cli"
	h "trains/helper/routeUtils"
=======
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	c "trains/cli"
>>>>>>> 3-gracefull-shutdown
	p "trains/pathfinder"
	s "trains/scheduler"
)

func main() {
<<<<<<< HEAD
	//go run . -feature stations.map waterloo euston 5
	//go run . -feature jungle-desert.map jungle desert 5
	//go run . -feature smallAndLarge.map small large 9
=======
	// go run . -feature file start end 7
>>>>>>> 3-gracefull-shutdown
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

	//res := p.BigFuckingSearch(appData)
	res := p.BuildFlowGraph(&appData)

	maxFlow := res.Graph.MaxFlow(res.StationToID[appData.StartingStation], res.StationToID[appData.EndingStation])
	paths, err := res.ExtractPaths(maxFlow)

	for i, path := range paths {
		fmt.Printf("%d: ", i)
		for _, st := range path {
			fmt.Printf("%s ", st.Name)
		}
		fmt.Println()
	}
	fmt.Println("Loaded trains:", conf.TrainNumb)

<<<<<<< HEAD
=======
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
>>>>>>> 3-gracefull-shutdown
}
