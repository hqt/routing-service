package routing

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildGraph(t *testing.T) {
	t.Parallel()

	stations, err := ParseCSVToStations("../../etc/StationMap.csv")
	require.Nil(t, err)
	graph := BuildGraph(stations)
	require.NotNil(t, graph)

	// test idToNode
	require.Equal(t, len(stations), len(graph.idToNode))

	require.Equal(t, 3, len(graph.idToNode["EW21"].neighbors))
	var neighbors []string
	for _, node := range graph.idToNode["EW21"].neighbors {
		neighbors = append(neighbors, node.id())
	}
	assert.ElementsMatch(t, []string{"CC22", "EW20", "EW22"}, neighbors)

	neighbors = []string{}
	require.Equal(t, 1, len(graph.idToNode["EW1"].neighbors))
	for _, node := range graph.idToNode["EW1"].neighbors {
		neighbors = append(neighbors, node.id())
	}
	assert.ElementsMatch(t, []string{"EW2"}, neighbors)
}

func TestFindNodeByName(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		stations, err := ParseCSVToStations("../../etc/StationMap.csv")
		require.Nil(t, err)
		graph := BuildGraph(stations)
		ids := graph.FindNodeByName("Jurong East")
		require.Equal(t, 2, len(ids))
		assert.ElementsMatch(t, []string{"NS1", "EW24"}, ids)
	})
}

func TestPrintInstructions(t *testing.T) {
	t.Parallel()

	stations, err := ParseCSVToStations("../../etc/StationMap.csv")
	require.Nil(t, err)
	graph := BuildGraph(stations)
	routes := []string{"CC21", "CC20", "CC19", "DT9", "DT10", "DT11", "DT12", "DT13", "DT14"}
	instructions := graph.PrintInstructions(routes)
	require.Equal(t, len(routes)-1, len(instructions))
}
