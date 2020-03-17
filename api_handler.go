package zendesk

import (
	"encoding/json"
	"net/http"
	"time"
)

func (app *App) simpleRoute(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var params map[string]interface{}
	err := dec.Decode(&params)
	if err != nil {
		SendJSON(w, http.StatusBadRequest, nil)
		return
	}

	from, ok := params["from"].(string)
	if !ok || len(from) == 0 {
		SendJSON(w, http.StatusBadRequest, nil)
		return
	}

	to := params["to"].(string)
	if !ok || len(to) == 0 {
		SendJSON(w, http.StatusBadRequest, nil)
		return
	}

	paths, ok := app.graph.FindRoutes(from, to, time.Now())
	if !ok {
		SendJSON(w, http.StatusOK, map[string]interface{}{
			"verdict": "not_found",
		})
		return
	}

	instructions := app.graph.PrintInstructions(paths)
	body := map[string]interface{}{
		"start_at":     time.Now(),
		"verdict":      "success",
		"paths":        paths,
		"instructions": instructions,
	}

	SendJSON(w, http.StatusOK, body)
}
