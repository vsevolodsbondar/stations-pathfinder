package pathfinder

import (
	"errors"
	"fmt"
	"math"
	m "trains/models"
)

// Rail is one directed rail in the flow graph.
type Rail struct {
	To       int
	Capacity int
	Flow     int
	Reverse  int
}

// FlowGraph stores all rails used by the max flow algorithm.
type FlowGraph struct {
	Rails [][]Rail
}

// FlowNetwork connects the railway map with the flow graph.
type FlowNetwork struct {
	Graph       *FlowGraph
	StationToID map[*m.Station]int
	Stations    []*m.Station
	Start       *m.Station
	End         *m.Station
}

// Parent stores the previous node while searching for a path.
type Parent struct {
	From int
	Rail *Rail
}

// BuildFlowGraph converts the railway network into a flow graph.
// Each station is split into an input and an output node so that
// only one train can pass through intermediate stations.
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
		Stations:    data.NetworkMap,
		Start:       data.StartingStation,
		End:         data.EndingStation,
	}
}

// RemainingCapacity returns how much flow can still pass through this rail.
func (r Rail) RemainingCapacity() int {
	return r.Capacity - r.Flow
}

func (g *FlowNetwork) Print() {
	for i, rails := range g.Graph.Rails {
		fmt.Printf("%d:\n", i)
		for _, rail := range rails {
			fmt.Printf("->%d cap=%d flow=%d rev=%d\n", rail.To, rail.Capacity, rail.Flow, rail.Reverse)
		}
	}
}

// AddRail adds a rail and its reverse rail to the graph.
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

// BFS finds an augmenting path using breadth-first search.
func (g *FlowGraph) BFS(start, end int) ([]Parent, bool) {
	if len(g.Rails) == 0 {
		return nil, false
	}
	parents := make([]Parent, len(g.Rails))
	visited := make([]bool, len(g.Rails))
	visited[start] = true
	queue := []int{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for i := range g.Rails[current] {
			rail := &g.Rails[current][i]
			if rail.RemainingCapacity() <= 0 {
				continue
			}
			next := rail.To
			if visited[next] {
				continue
			}
			visited[next] = true
			parents[next] = Parent{
				From: current,
				Rail: rail,
			}
			if next == end {
				return parents, true
			}
			queue = append(queue, next)
		}
	}
	return parents, false
}

// MaxFlow finds the maximum number of independent paths.
func (g *FlowGraph) MaxFlow(start, end int) int {
	maxFlow := 0
	for {
		parents, found := g.BFS(start, end)
		if !found {
			break
		}
		pathFlow := math.MaxInt
		current := end
		for current != start {
			p := parents[current]
			if p.Rail.RemainingCapacity() < pathFlow {
				pathFlow = p.Rail.RemainingCapacity()
			}
			current = p.From
		}
		current = end
		for current != start {
			p := parents[current]
			p.Rail.Flow += pathFlow
			reverse := &g.Rails[p.Rail.To][p.Rail.Reverse]
			reverse.Flow -= pathFlow
			current = p.From
		}
		maxFlow += pathFlow
	}
	return maxFlow
}

// ExtractPaths builds railway paths from the computed flow.
func (n *FlowNetwork) ExtractPaths(pathCount int) ([][]*m.Station, error) {
	paths := make([][]*m.Station, 0, pathCount)
	startIn := n.StationToID[n.Start]
	endIn := n.StationToID[n.End]
	for i := 0; i < pathCount; i++ {
		path := []*m.Station{}
		current := startIn
		for current != endIn {
			if current%2 == 0 {
				path = append(path, n.Stations[current/2])
			}
			found := false
			for j := range n.Graph.Rails[current] {
				rail := &n.Graph.Rails[current][j]
				if rail.Flow <= 0 {
					continue
				}
				rail.Flow--
				reverse := &n.Graph.Rails[rail.To][rail.Reverse]
				reverse.Flow++
				current = rail.To
				found = true
				break
			}
			if !found {
				return nil, errors.New("failed to extract path")
			}
		}
		path = append(path, n.End)
		paths = append(paths, path)
	}
	return paths, nil
}
