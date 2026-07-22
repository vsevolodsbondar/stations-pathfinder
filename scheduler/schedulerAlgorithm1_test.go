package scheduler

import (
	"testing"
	m "trains/models"
)

func BenchmarkScheduleTrains(b *testing.B) {

	routes := []m.Route{
		{
			ID:         1,
			Route:      []string{"jungle", "grasslands", "suburbs", "clouds", "wetlands", "desert"},
			RouteScore: 6,
		},
		{
			ID:         2,
			Route:      []string{"jungle", "grasslands", "suburbs", "clouds", "wetlands", "mountain", "treetop", "desert"},
			RouteScore: 8,
		},
		{
			ID:         3,
			Route:      []string{"jungle", "grasslands", "suburbs", "clouds", "wetlands", "mountain", "farms", "downtown", "metropolis", "industrial", "desert"},
			RouteScore: 1,
		},
		{
			ID:         4,
			Route:      []string{"jungle", "farms", "downtown", "metropolis", "industrial", "desert"},
			RouteScore: 6,
		},
		// {
		// 	ID:         5,
		// 	Route:      []string{"jungle", "farms", "mountain", "treetop", "desert"},
		// 	RouteScore: 5,
		// },
	}
	trainNumb := 10000

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rs := make([]m.Route, len(routes))
		copy(rs, routes)

		MoveTrains(rs, trainNumb)
	}
}
