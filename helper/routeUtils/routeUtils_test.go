package routeutils

import (
	"fmt"
	"testing"
	m "trains/models"
)

func mockRoutes() []m.Route {
	return []m.Route{
		{
			ID:             1,
			Route:          []string{"jungle", "farms", "mountain", "wetlands", "desert"},
			Distance:       36.76453843232368,
			CrossingRoutes: map[int]struct{}{},
		},
		{
			ID:             2,
			Route:          []string{"jungle", "farms", "mountain", "treetop", "desert"},
			Distance:       39.65869449637143,
			CrossingRoutes: map[int]struct{}{},
		},
		{
			ID:             3,
			Route:          []string{"jungle", "green_belt", "village", "mountain", "wetlands", "desert"},
			Distance:       52.91961844821559,
			CrossingRoutes: map[int]struct{}{},
		},
		{
			ID:             4,
			Route:          []string{"jungle", "green_belt", "village", "mountain", "treetop", "desert"},
			Distance:       55.813774512263336,
			CrossingRoutes: map[int]struct{}{},
		},
		{
			ID:             5,
			Route:          []string{"jungle", "farms", "downtown", "metropolis", "industrial", "desert"},
			Distance:       56.262188101751256,
			CrossingRoutes: map[int]struct{}{},
		},
		{
			ID:             6,
			Route:          []string{"jungle", "grasslands", "suburbs", "clouds", "wetlands", "desert"},
			Distance:       58.051833271472525,
			CrossingRoutes: map[int]struct{}{},
		},
		{
			ID:             7,
			Route:          []string{"jungle", "grasslands", "suburbs", "clouds", "wetlands", "mountain", "treetop", "desert"},
			Distance:       77.07050483211736,
			CrossingRoutes: map[int]struct{}{},
		},
		{
			ID:             8,
			Route:          []string{"jungle", "green_belt", "village", "mountain", "farms", "downtown", "metropolis", "industrial", "desert"},
			Distance:       85.06637875831666,
			CrossingRoutes: map[int]struct{}{},
		},
		{
			ID:             9,
			Route:          []string{"jungle", "grasslands", "suburbs", "clouds", "wetlands", "mountain", "farms", "downtown", "metropolis", "industrial", "desert"},
			Distance:       106.32310907817072,
			CrossingRoutes: map[int]struct{}{},
		},
	}
}

func TestShouldAddCrossingRoutesCorrectly(t *testing.T) {
	data := mockRoutes()

	findCrossingRoutes(data)

	for _, v := range data {
		fmt.Println(v)
	}

	if len(data[0].CrossingRoutes) == 0 {
		t.Errorf("Should have crossing routes data.")
	}
}
