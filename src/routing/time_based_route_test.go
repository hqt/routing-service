package routing

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTimeBasedRouting(t *testing.T) {
	t.Parallel()

	stations, err := ParseCSVToStations("../../etc/StationMap.csv")
	require.Nil(t, err)
	graph := BuildGraph(stations)

	source := "Boon Lay"
	destination := "Little India"
	startTimeStr := "2019-01-31T16:00"
	layout := "2006-01-02T15:04"
	startTime, err := time.Parse(layout, startTimeStr)
	require.Nil(t, err)

	expectedRoutes := []string{"EW27", "EW26", "EW25", "EW24", "EW23", "EW22", "EW21", "CC22", "CC21", "CC20", "CC19", "DT9", "DT10", "DT11", "DT12"}
	routes, cost, found := graph.FindRoutesWithConstraints(source, destination, startTime)
	require.True(t, found)
	require.Equal(t, 134, cost)
	require.Equal(t, expectedRoutes, routes)
}
