package zendesk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseCSVToRoutes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name            string
		Path            string
		ExpectedSuccess bool
		TotalRecords    int
	}{
		{"file_not_found", "not_found.csv", false, 0},
		{"missing_field", "etc/csv_test/missing_field.csv", false, 0},
		{"wrong_date_format", "etc/csv_test/wrong_date_format.csv", false, 0},
		{"invalid_station_format", "etc/csv_test/invalid_station_format.csv", false, 0},
		{"invalid_station_line_format", "etc/csv_test/invalid_station_line_format.csv", false, 0},
		{"full_data_file", "etc/StationMap.csv", true, 166},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			routes, err := ParseCSVToStations(tc.Path)
			if tc.ExpectedSuccess {
				require.Nil(t, err)
				require.Equal(t, tc.TotalRecords, len(routes))
			} else {
				require.NotNil(t, err)
			}
		})
	}
}
