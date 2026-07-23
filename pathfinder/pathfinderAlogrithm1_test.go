package pathfinder

import (
	"fmt"
	"testing"
	m "trains/models"
)

func generateLargeAppData(size int) m.AppData {
	if size <= 0 {
		return m.AppData{}
	}

	//limitations by the task
	if size > 10000 {
		return m.AppData{}
	}

	stations := make([]*m.Station, size+1)

	for i := 0; i <= size; i++ {
		stations[i] = &m.Station{
			Name:   fmt.Sprintf("%d", i+1),
			X_axis: i,
			Y_axis: 0,
		}
	}

	for i := 0; i < size-1; i++ {
		stations[i].Connections = append(stations[i].Connections, stations[i+1])
		stations[i+1].Connections = append(stations[i+1].Connections, stations[i])
	}

	return m.AppData{
		NetworkMap:      stations,
		StartingStation: stations[0],
		EndingStation:   stations[size-1],
		TrainNumb:       1,
	}
}

// nodes = 2^depth - 1
func generateBinaryGraph(depth int) m.AppData {
	if depth <= 0 {
		return m.AppData{}
	}

	if depth > 14 {
		return m.AppData{}
	}

	// Number of nodes in a full binary tree.
	size := (1 << depth) - 1

	stations := make([]*m.Station, size)

	// Create stations.
	for i := range stations {
		stations[i] = &m.Station{
			Name:   fmt.Sprintf("%d", i+1),
			X_axis: i,
			Y_axis: depth,
		}
	}

	// Connect parent <-> children.
	for i := 0; i < size; i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < size {
			stations[i].Connections = append(stations[i].Connections, stations[left])
			stations[left].Connections = append(stations[left].Connections, stations[i])
		}

		if right < size {
			stations[i].Connections = append(stations[i].Connections, stations[right])
			stations[right].Connections = append(stations[right].Connections, stations[i])
		}
	}

	// Choose the leftmost leaf as the destination.
	end := stations[(1<<(depth-1))-1]

	return m.AppData{
		NetworkMap:      stations,
		StartingStation: stations[0],
		EndingStation:   end,
		TrainNumb:       1,
	}
}

func BenchmarkFindPathsWithBinaryData(b *testing.B) {
	//depth should be notthing more than 14, as for 14 there will be already 16 383 nodes
	g := generateBinaryGraph(14)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DFSRangedRouteSets(g)
	}
}

func BenchmarkFindPathsWithLargeData(b *testing.B) {
	g := generateLargeAppData(10000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DFSRangedRouteSets(g)
	}
}
