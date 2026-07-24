package service

import (
	"context"
	"fmt"
	"log"
	r "trains/helper/routeUtils"
	m "trains/models"
	p "trains/pathfinder"
	s "trains/scheduler"
)

func Algorithm1Runner(ctx context.Context, appData m.AppData) error {
	log.Println("Running Seva's algorithm.")
	log.Println("Starting to search for paths.")
	routeSets, err := p.DFSRangedRouteSets(appData)
	if err != nil {
		return err
	}
	best := r.BestRoutes(routeSets, appData.TrainNumb)
	log.Println("Found best paths.")

	select {
	case <-ctx.Done():
		log.Println("Scheduler is stopping the work...")
		return nil

	default:
	}

	log.Println("Starting to move trains.")
	s.MoveTrains(best, appData.TrainNumb)
	log.Println("Trains finished moving.")

	return nil
}

func Algorithm2Runner(ctx context.Context, appData m.AppData) error {
	log.Println("Running Anatolii's algorithm.")
	log.Println("Starting to search for paths.")
	res := p.BuildFlowGraph(&appData)
	startID := res.StationToID[appData.StartingStation]
	endID := res.StationToID[appData.EndingStation]
	trains := appData.TrainNumb

	maxFlow := res.Graph.MaxFlow(startID, endID)
	paths, err := res.ExtractPaths(maxFlow)
	if err != nil {
		return err
	}

	log.Println("Found best paths.")

	select {
	case <-ctx.Done():
		log.Println("Scheduler is stopping the work...")
		return nil

	default:
	}

	log.Println("Starting to move trains.")
	s.DistributeTrains(paths, trains)
	lines := s.Schedule(paths, trains)

	for _, line := range lines {
		fmt.Println(line)
	}

	log.Println("Trains finished moving.")

	return nil
}
