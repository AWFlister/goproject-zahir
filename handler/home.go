package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) HomeIndex(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	body := map[string]string{"message": "Hello Zahir!"}
	encoder.Encode(body)
}
