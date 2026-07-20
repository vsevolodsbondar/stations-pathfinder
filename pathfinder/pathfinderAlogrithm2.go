package pathfinder

import (
	"fmt"
	m "trains/models"
)

type Rail struct {
	To       int
	Capacity int
	Flow     int
	Reverse  int
}

type FlowGraph struct {
	Rails [][]Rail
}

type FlowNetwork struct {
	Graph       *FlowGraph
	StationToID map[*m.Station]int
}

type Parent struct {
	From int
	Rail *Rail
}

func BuildFlowGraph(data *m.AppData) *FlowNetwork {
	graph := &FlowGraph{
		Rails: make([][]Rail, len(data.NetworkMap)*2),
	}
	stationToID := make(map[*m.Station]int)

	id := 0
	for _, station := range data.NetworkMap {
		stationToID[station] = id
		id += 2
	}
	for _, station := range data.NetworkMap {

		in := stationToID[station]
		out := in + 1
		capacity := 1

		if station == data.StartingStation || station == data.EndingStation {
			capacity = data.TrainNumb
		}

		graph.AddRail(in, out, capacity)
	}
	for _, station := range data.NetworkMap {

		fromOut := stationToID[station] + 1

		for _, next := range station.Connections {

			toIn := stationToID[next]

			graph.AddRail(fromOut, toIn, 1)
		}
	}
	return &FlowNetwork{
		Graph:       graph,
		StationToID: stationToID,
	}
}

func (r Rail) RemainingCapacity() int {
	return r.Capacity - r.Flow
}

func (g *FlowGraph) Print() {
	for i, rails := range g.Rails {
		fmt.Printf("%d:\n", i)
		for _, rail := range rails {
			fmt.Printf("->%d cap=%d flow=%d rev=%d\n", rail.To, rail.Capacity, rail.Flow, rail.Reverse)
		}
	}
}

func (g *FlowGraph) AddRail(from, to, capacity int) {
	forwardIdx := len(g.Rails[from])
	reverseIdx := len(g.Rails[to])
	forward := Rail{
		To:       to,
		Capacity: capacity,
		Flow:     0,
		Reverse:  reverseIdx,
	}

	reverse := Rail{
		To:       from,
		Capacity: 0,
		Flow:     0,
		Reverse:  forwardIdx,
	}
	g.Rails[from] = append(g.Rails[from], forward)
	g.Rails[to] = append(g.Rails[to], reverse)
}

func (g *FlowGraph) BFS(start, end int) ([]Parent, bool) {
	var parents = make([]Parent, len(g.Rails))
	visited := make([]bool, len(g.Rails))
	queue := []int{start}
	current := queue[0]
	queue = queue[1:]

}
