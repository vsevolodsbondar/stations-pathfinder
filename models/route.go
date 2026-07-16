package models

type Route struct {
	Route          []string
	Distance       float64
	CrossingRoutes []Route
}
