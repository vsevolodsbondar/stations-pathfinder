package scheduler

import (
<<<<<<< HEAD
	"fmt"
	h "trains/helper/schedulingUtils"
	m "trains/models"
)

func MoveTrains(routes []m.Route, trainsNumb int) {
	trains := h.TrainMaker(routes[0].Route[0], trainsNumb)
	AssignRouteToTrain(trains, routes)

	freeStations := map[string]bool{}

	// Only intermediate stations are tracked.
	for _, r := range routes {
		for _, s := range r.Route[1 : len(r.Route)-1] {
			freeStations[s] = true
		}
	}

	for !allFinished(trains) {
		turnString := ""

		for i := range trains {
			if trains[i].Finished {
				continue
			}

			nextStation := trains[i].Route.Route[trains[i].Turn+1]
			endStation := trains[i].Route.Route[len(trains[i].Route.Route)-1]

			if nextStation != endStation && !freeStations[nextStation] {
				continue
			}

			prevStation := trains[i].CurrStation

			// freeing prev station
			_, exists := freeStations[prevStation]
			if exists {
				freeStations[prevStation] = true
			}

			// occupying next station if it's not the last one
			if nextStation != endStation {
				freeStations[nextStation] = false
			}

			_, info := trains[i].Move()
			turnString += info
		}

		if turnString != "" {
			fmt.Println(turnString)
		}
	}
}

func AssignRouteToTrain(trains []m.Train, routes []m.Route) {
	for i := range trains {
		idx := bestTrainRoute(routes)
		trains[i].Route = routes[idx]
	}
}

func bestTrainRoute(routes []m.Route) int {
	best := 0

	for i := 1; i < len(routes); i++ {
		if routes[i].RouteScore < routes[best].RouteScore {
			best = i
		} else if routes[i].RouteScore == routes[best].RouteScore &&
			routes[i].Distance < routes[best].Distance {
			best = i
		}
	}

	routes[best].RouteScore++

	return best
}

func allFinished(trains []m.Train) bool {
	for _, t := range trains {
		if !t.Finished {
			return false
		}
	}
	return true
}
=======
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
>>>>>>> 3-gracefull-shutdown
