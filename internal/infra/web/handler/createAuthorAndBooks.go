package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelsouzaribeiro/uow-golang/internal/dto"
)

func (h *Handler) CreateAuthorAndBooks(w http.ResponseWriter, r *http.Request) {
	var payload dto.AuthorBooks

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.usecase.CreateAuthorAndBook(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
