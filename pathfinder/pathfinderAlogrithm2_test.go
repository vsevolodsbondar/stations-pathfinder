package pathfinder

import (
	"fmt"
	"testing"
	m "trains/models"
)

func generateParallelGraph(paths int) m.AppData {
	stations := []*m.Station{}

	start := &m.Station{Name: "start"}
	end := &m.Station{Name: "end"}

	stations = append(stations, start)

	for i := 0; i < paths; i++ {
		mid := &m.Station{
			Name: fmt.Sprintf("mid%d", i),
		}

		start.Connections = append(start.Connections, mid)
		mid.Connections = append(mid.Connections, start)

		mid.Connections = append(mid.Connections, end)

		stations = append(stations, mid)
	}

	stations = append(stations, end)

	return m.AppData{
		NetworkMap:      stations,
		StartingStation: start,
		EndingStation:   end,
		TrainNumb:       paths,
	}
}

func BenchmarkPathFinderLinear(b *testing.B) {
	data := generateLargeAppData(10000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		network := BuildFlowGraph(&data)

		start := network.StationToID[data.StartingStation]
		end := network.StationToID[data.EndingStation]

		count := network.Graph.MaxFlow(start, end)

		_, _ = network.ExtractPaths(count)
	}
}

func BenchmarkPathFinderBinary(b *testing.B) {
	data := generateBinaryGraph(14)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		network := BuildFlowGraph(&data)

		start := network.StationToID[data.StartingStation]
		end := network.StationToID[data.EndingStation]

		count := network.Graph.MaxFlow(start, end)

		_, _ = network.ExtractPaths(count)
	}
}

func BenchmarkPathFinderParallel(b *testing.B) {
	data := generateParallelGraph(1000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		network := BuildFlowGraph(&data)

		start := network.StationToID[data.StartingStation]
		end := network.StationToID[data.EndingStation]

		count := network.Graph.MaxFlow(start, end)

		_, _ = network.ExtractPaths(count)
	}
}
