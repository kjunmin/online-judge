package common

import (
	"encoding/json"
	"net/http"
)

func WriteJSONToHTTPResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func ReadJSONFromHTTPRequest(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteErrorToHTTPResponse(w http.ResponseWriter, status int, message string) {
	WriteJSONToHTTPResponse(w, status, map[string]string{"error": message})
}
