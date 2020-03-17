package routing

import (
	"github.com/hqt/routing-service/src/datastructure"
	"math"
	"time"
)

// INF defines infinity number and without overflow when calculating
const INF = math.MaxInt32 / 10

// FindRoutesWithConstraints finds routes with constraints using Dijkstra algorithm
func (g Graph) FindRoutesWithConstraints(start string, end string, queryTime time.Time) ([]string, int, bool) {
	if start == end {
		return []string{}, 0, true
	}

	// stored all visited node
	visited := map[string]bool{}
	// for tracing back the result
	prev := map[string]string{}
	// for tracking the cost of nodes
	costs := map[string]int{}

	allSources := g.FindNodeByName(start)
	allDestinations := g.FindNodeByName(end)

	// assign all nodes with costs is INF
	for id := range g.idToNode {
		costs[id] = INF
	}

	heap := datastructure.NewPriorityQueue()
	for _, source := range allSources {
		if !g.idToNode[source].station.IsOpen(queryTime) {
			continue
		}
		heap.Push(datastructure.NewHeapNode(source, 0))
		costs[source] = 0
		visited[source] = true
	}

	for heap.Len() > 0 {
		node := heap.Pop()
		stationID := node.Key

		neighbors := g.idToNode[stationID].neighbors
		for _, neighbor := range neighbors {
			neighborID := neighbor.id()
			if visited[neighborID] {
				continue
			}
			if !g.idToNode[neighborID].station.IsOpen(queryTime) {
				continue
			}

			// we must update again current time. because the time we go to the stationID, we might go to different time price
			currentTime := queryTime.Add(time.Minute * time.Duration(costs[stationID]))
			weight := getTimeBetweenStation(g.idToNode[stationID].station, g.idToNode[neighborID].station, currentTime)

			updatedWeight := costs[stationID] + weight
			if updatedWeight < costs[neighborID] {
				costs[neighborID] = updatedWeight
				prev[neighborID] = stationID
				heap.Push(datastructure.NewHeapNode(neighborID, updatedWeight))
			}

			visited[neighborID] = true
		}
	}

	foundStationID := ""
	min := INF
	for _, dest := range allDestinations {
		if costs[dest] < min {
			min = costs[dest]
			foundStationID = dest
		}
	}

	if foundStationID == "" {
		return nil, 0, false
	}

	return rebuildSolution(prev, foundStationID), costs[foundStationID], true
}
