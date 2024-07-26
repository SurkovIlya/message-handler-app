package server

import (
	"encoding/json"
	"mime"
	"net/http"
)

func (s *Server) ReceivingMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)

		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var rp Message

	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = s.Messager.Receiving(rp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(http.StatusText(200))
}

func (s *Server) GetStatistics(w http.ResponseWriter, r *http.Request) {

}
