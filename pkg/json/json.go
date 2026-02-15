package json

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, s int, d any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(d)
}

func ReadRequest(r *http.Request, d any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(d)
}
