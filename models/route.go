package models

type Route struct {
	ID             int
	Route          []string
	Distance       float64
	CrossingRoutes map[int]struct{}
}
