package server

import (
	"encoding/json"
	"net/http"
	"time"
)

const dateTimeLayout = "2006-01-02T15:04"

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

func (app *App) advancedRoute(w http.ResponseWriter, r *http.Request) {
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

	dateStr := params["start_time"].(string)
	if !ok || len(to) == 0 {
		SendJSON(w, http.StatusBadRequest, nil)
		return
	}

	startTime, err := time.Parse(dateTimeLayout, dateStr)
	if err != nil {
		SendJSON(w, http.StatusBadRequest, nil)
		return
	}

	paths, costs, ok := app.graph.FindRoutesWithConstraints(from, to, startTime)
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
		"minutes":      costs,
	}

	SendJSON(w, http.StatusOK, body)
}
