package routingservice

import (
	"fmt"
	"sort"
)

// Node presents a node in the graph. Each node is a station with all its neighbors
type Node struct {
	station   *Station
	neighbors []*Node
}

// Graph represents the whole graph
type Graph struct {
	idToNode       map[string]*Node
	lineToStations map[string][]*Node
}

// newNode creates a node from a station
func newNode(station *Station) *Node {
	return &Node{
		station: station,
	}
}

// addNeighbor adds a neighbor
func (n *Node) addNeighbor(neighbor *Node) {
	n.neighbors = append(n.neighbors, neighbor)
}

// id represents node's id
func (n *Node) id() string {
	return n.station.Code
}

// BuildGraph builds a whole graph from list of stations
func BuildGraph(stations []*Station) Graph {
	// build adjacency nodes share the same place but different code
	idToNodes := map[string]*Node{}
	n := len(stations)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			first := stations[i]
			second := stations[j]
			if idToNodes[first.Code] == nil {
				idToNodes[first.Code] = newNode(first)
			}
			if idToNodes[second.Code] == nil {
				idToNodes[second.Code] = newNode(second)
			}
			if first.Name == second.Name {
				idToNodes[first.Code].addNeighbor(idToNodes[second.Code])
				idToNodes[second.Code].addNeighbor(idToNodes[first.Code])
			}
		}
	}

	lineToStations := map[string][]*Node{}
	// collects all stations on the same line
	for _, node := range idToNodes {
		station := node.station
		if lineToStations[station.Line] == nil {
			lineToStations[station.Line] = []*Node{}
		}
		lineToStations[station.Line] = append(lineToStations[station.Line], node)
	}

	// sort again each stations on each line based on the number
	for _, nodes := range lineToStations {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].station.Order < nodes[j].station.Order
		})
	}

	// build adjacency nodes from 2 places on the same lines
	for _, nodes := range lineToStations {
		for i := 0; i < len(nodes)-1; i++ {
			first := nodes[i].station
			second := nodes[i+1].station
			idToNodes[first.Code].addNeighbor(idToNodes[second.Code])
			idToNodes[second.Code].addNeighbor(idToNodes[first.Code])
		}
	}

	return Graph{
		idToNode:       idToNodes,
		lineToStations: lineToStations,
	}
}

// FindNodeByName finds node by name. returns list all station ids that have the same name
func (g Graph) FindNodeByName(name string) []string {
	var res []string
	for _, node := range g.idToNode {
		if node.station.Name == name {
			res = append(res, node.station.Code)
		}
	}
	return res
}

// PrintInstructions prints all the instructions based on paths
func (g Graph) PrintInstructions(paths []string) []string {
	var res []string
	for i := 0; i < len(paths)-1; i++ {
		first := g.idToNode[paths[i]].station
		second := g.idToNode[paths[i+1]].station

		sentence := ""
		if first.Line == second.Line {
			sentence = fmt.Sprintf("Take %s line from %s to %s", first.Line, first.Name, second.Name)
		} else {
			sentence = fmt.Sprintf("Change from %s line to %s line", first.Line, second.Line)
		}
		res = append(res, sentence)
	}
	return res
}

func rebuildSolution(prev map[string]string, dest string) []string {
	var res []string
	for dest != "" {
		res = append(res, dest)
		dest = prev[dest]
	}

	// reverse array
	for i := 0; i < len(res)/2; i++ {
		tmp := res[i]
		res[i] = res[len(res)-1-i]
		res[len(res)-1-i] = tmp
	}

	return res
}
