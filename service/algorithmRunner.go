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
	routeSets, err := p.DFSRangedRouteSets(appData)
	if err != nil {
		return err
	}
	best := r.BestRoutes(routeSets, appData.TrainNumb)

	select {
	case <-ctx.Done():
		log.Println("Scheduler is stopping the work...")
		return nil

	default:
	}

	s.MoveTrains(best, appData.TrainNumb)
	log.Println("Trains finished moving.")

	return nil
}
