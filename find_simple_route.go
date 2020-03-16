package zendesk

// FindRoutes returns a path from source to dest using BFS
func (g *Graph) FindRoutes(start string, end string) ([]string, bool) {
	if start == end {
		return []string{}, true
	}

	// stored all visited node
	visited := map[string]bool{}
	// for tracing back the result
	prev := map[string]string{}

	allSources := g.FindNodeByName(start)
	allDestinations := g.FindNodeByName(end)

	queue := NewFIFO()
	for _, source := range allSources {
		queue.Push(source)
		visited[source] = true
		prev[source] = ""
	}

	foundStationID := ""
	for queue.Len() > 0 {
		stationID := queue.Front().(string)
		if stringsContain(allDestinations, stationID) {
			foundStationID = stationID
			break
		}

		neighbors := g.idToNode[stationID].neighbors
		for _, neighbor := range neighbors {
			if visited[neighbor.id()] {
				continue
			}
			visited[neighbor.id()] = true
			prev[neighbor.id()] = stationID
			queue.Push(neighbor.id())
		}
	}

	if foundStationID == "" {
		return nil, false
	}

	return rebuildSolution(prev, foundStationID), true
}
