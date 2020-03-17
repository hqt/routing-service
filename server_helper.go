package routingservice

import (
	"encoding/json"
	"log"
	"net/http"
)

// define constants
const (
	HeaderContentType = "Content-Type"
	ContentTypeJSON   = "application/json"
)

// SendJSON sends json back to client
func SendJSON(w http.ResponseWriter, statusCode int, data map[string]interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(statusCode)

	body, err := json.Marshal(data)
	if err != nil {
		log.Println("cannot marshal response data:", err)
	}

	_, err = w.Write(body)
	if err != nil {
		log.Println("cannot write response body:", err)
	}
}
