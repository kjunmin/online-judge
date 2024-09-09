package runner

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/run", h.handleCodeExecution)
}

func (h *Handler) handleCodeExecution(w http.ResponseWriter, r *http.Request) {
	return
}
