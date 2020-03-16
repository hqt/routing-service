package zendesk

import (
	"github.com/stretchr/testify/require"
	"testing"
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
		routes, found := graph.FindRoutes(source, destination)
		require.True(t, found)
		require.Equal(t, expectedRoutes, routes)
	})

	t.Run("source is same with dest", func(t *testing.T) {
		t.Parallel()
		source := "Holland Village"
		dest := "Holland Village"
		routes, found := graph.FindRoutes(source, dest)
		require.True(t, found)
		require.Equal(t, 0, len(routes))
	})

	t.Run("not_found", func(t *testing.T) {
		source := "Holland Village"
		destination := "Bugis"
		expectedRoutes := []string{"CC21", "CC20", "CC19", "DT9", "DT10", "DT11", "DT12", "DT13", "DT14"}
		routes, found := graph.FindRoutes(source, destination)
		require.True(t, found)
		require.Equal(t, expectedRoutes, routes)
	})
}
