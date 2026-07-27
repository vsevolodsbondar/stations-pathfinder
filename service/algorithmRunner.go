package service

import (
	"context"
	"fmt"
	"log"
	"time"
	r "trains/helper/routeUtils"
	m "trains/models"
	p "trains/pathfinder"
	s "trains/scheduler"
)

func Algorithm1Runner(ctx context.Context, appData m.AppData) error {
	log.Println("Running Seva's algorithm.")
	log.Println("Starting to search for paths.")

	routeDone := make(chan struct{})
	var best []m.Route
	go func() error {
		routeSets, err := p.DFSRangedRouteSets(appData)
		if err != nil {
			return err
		}
		best = r.BestRoutes(routeSets, appData.TrainNumb)
		log.Println("Found best paths.")
		close(routeDone)

		return nil
	}()

	//waiting for either completion or timeout
	select {
	case <-routeDone:
	case <-ctx.Done():
		fmt.Println("Shutdown performed before found all paths.")
		return nil
	case <-time.After(10 * time.Second):
		return fmt.Errorf("Paths can't be found in appropriate time.")
	}

	moveDone := make(chan struct{})
	go func() {
		log.Println("Starting to move trains.")
		s.MoveTrains(best, appData.TrainNumb)
		close(moveDone)
	}()

	select {
	case <-moveDone:
		log.Println("Trains finished moving.")

	case <-time.After(5 * time.Second):
		log.Println("Trains moving timed out.")
	}

	return nil
}

func Algorithm2Runner(ctx context.Context, appData m.AppData) error {
	log.Println("Running Anatolii's algorithm.")
	log.Println("Starting to search for paths.")

	trains := appData.TrainNumb
	routeDone := make(chan struct{})
	var paths [][]*m.Station

	go func() error {
		res := p.BuildFlowGraph(&appData)
		startID := res.StationToID[appData.StartingStation]
		endID := res.StationToID[appData.EndingStation]

		maxFlow := res.Graph.MaxFlow(startID, endID)
		var err error
		paths, err = res.ExtractPaths(maxFlow)
		if err != nil {
			return err
		}
		log.Println("Found best paths.")
		close(routeDone)

		return nil
	}()

	select {
	case <-routeDone:
	case <-ctx.Done():
		fmt.Println("Shutdown performed before found all paths.")
		return nil
	case <-time.After(10 * time.Second):
		return fmt.Errorf("Paths can't be found in appropriate time.")
	}

	moveDone := make(chan struct{})
	go func() {
		log.Println("Starting to move trains.")
		s.DistributeTrains(paths, trains)
		lines := s.Schedule(paths, trains)

		for _, line := range lines {
			fmt.Println(line)
		}
		close(moveDone)
	}()

	select {
	case <-moveDone:
		log.Println("Trains finished moving.")

	case <-time.After(5 * time.Second):
		log.Println("Trains moving timed out.")
	}

	return nil
}
