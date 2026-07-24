package scheduler

import (
	"testing"

	m "trains/models"
)

func BenchmarkSchedule(b *testing.B) {
	makePath := func(names ...string) []*m.Station {
		path := make([]*m.Station, len(names))
		for i, name := range names {
			path[i] = &m.Station{Name: name}
		}
		return path
	}

	paths := [][]*m.Station{
		makePath("jungle", "grasslands", "suburbs", "clouds", "wetlands", "desert"),
		makePath("jungle", "grasslands", "suburbs", "clouds", "wetlands", "mountain", "treetop", "desert"),
		makePath("jungle", "grasslands", "suburbs", "clouds", "wetlands", "mountain", "farms", "downtown", "metropolis", "industrial", "desert"),
		makePath("jungle", "farms", "downtown", "metropolis", "industrial", "desert"),
	}

	trainCount := 10000

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Schedule(paths, trainCount)
	}
}
