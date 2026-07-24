package service

import (
	"context"
	"log"
	r "trains/helper/routeUtils"
	m "trains/models"
	p "trains/pathfinder"
	s "trains/scheduler"
)

func AlgorithmRunner(ctx context.Context, appData m.AppData) error {
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
