package server

import (
	"encoding/json"
	"mime"
	"net/http"

	"github.com/SurkovIlya/message-handler-app/internal/model"
)

const maxMsgSize = 10e2

// @Summary Receiving messages
// @Description Receiving messages for further processing
// @ID receiving-messages
// @Accept  json
// @Produce  json
// @Param input body model.Message true "messages"
// @Success 200 {string} string "OK"
// @Failure 400  {string} error "the maximum allowed message size has been exceeded"
// @Router /v1/receiving [post]
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

	var rp model.Message

	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if len(rp.Value) > maxMsgSize {
		http.Error(w, "the maximum allowed message size has been exceeded", http.StatusBadRequest)

		return
	}

	err = s.Messager.Receive(rp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(http.StatusText(200))
}

// @Summary      Get statistics
// @Description  Get statistics on processed messages
// @ID get-statistics
// @Produce      json
// @Success      200  {object}  model.Statistic
// @Failure      500  {string} error err
// @Router       /v1/getstatistics [get]
func (s *Server) GetStatistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stat, err := s.SC.GetStat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(stat)
}
