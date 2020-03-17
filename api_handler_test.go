package routingservice

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestSimpleRouting(t *testing.T) {
	t.Parallel()

	from := "Holland Village"
	to := "Bugis"

	testCases := []struct {
		Name           string
		From           string
		To             string
		ExpectedStatus int
	}{
		{"empty_from", "", to, http.StatusBadRequest},
		{"empty_to", from, "", http.StatusBadRequest},
		{"success_not_found", "abc", "xyz", http.StatusOK},
		{"success", from, to, http.StatusOK},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			app, err := NewApp()
			require.Nil(t, err)
			server, err := app.StartTestServer()
			require.Nil(t, err)
			defer server.Close()

			values := map[string]string{
				"from": tc.From,
				"to":   tc.To,
			}

			url := server.URL + "/api/simple_route"
			jsonValue, _ := json.Marshal(values)
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
			require.Nil(t, err)
			require.NotNil(t, resp)
			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)

			decoder := json.NewDecoder(resp.Body)
			var data map[string]interface{}
			err = decoder.Decode(&data)
			require.Nil(t, err)

			if tc.Name == "success_not_found" {
				verdict := data["verdict"]
				require.Equal(t, "not_found", verdict)
			} else if tc.Name == "success" {
				instructions, ok := data["instructions"].([]interface{})
				require.True(t, ok)

				paths, ok := data["paths"].([]interface{})
				require.True(t, ok)

				require.Equal(t, len(paths)-1, len(instructions))
				expectedRoutes := []interface{}{"CC21", "CC20", "CC19", "DT9", "DT10", "DT11", "DT12", "DT13", "DT14"}
				require.Equal(t, expectedRoutes, paths)
			}
		})
	}
}
