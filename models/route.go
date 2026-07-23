package models

import "fmt"

type Route struct {
	ID             int
	Route          []string
	Distance       float64
	CrossingRoutes map[int]struct{}

	AssignedTrains int
	RouteScore     int
}

func (r Route) PrintRoute() {
	formatted := fmt.Sprintf("Route ID: %d, Distance: %.2f, %v", r.ID, r.Distance, r.Route)
	fmt.Println(formatted)
}
