package routingservice

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestFindRoutes(t *testing.T) {
	t.Parallel()

	stations, err := ParseCSVToStations("etc/StationMap.csv")
	require.Nil(t, err)
	graph := BuildGraph(stations)

	t.Run("success", func(t *testing.T) {
		source := "Holland Village"
		destination := "Bugis"
		expectedRoutes := []string{"CC21", "CC20", "CC19", "DT9", "DT10", "DT11", "DT12", "DT13", "DT14"}
		routes, found := graph.FindRoutes(source, destination, time.Now())
		require.True(t, found)
		require.Equal(t, expectedRoutes, routes)
	})

	t.Run("source is same with dest", func(t *testing.T) {
		t.Parallel()
		source := "Holland Village"
		dest := "Holland Village"
		routes, found := graph.FindRoutes(source, dest, time.Now())
		require.True(t, found)
		require.Equal(t, 0, len(routes))
	})

	t.Run("not_found_because_no_opening_stations", func(t *testing.T) {
		source := "Holland Village"
		destination := "Bugis"
		pastTime := time.Now().AddDate(-100, 0, 0)
		routes, found := graph.FindRoutes(source, destination, pastTime)
		require.False(t, found)
		require.Empty(t, routes)
	})
}
