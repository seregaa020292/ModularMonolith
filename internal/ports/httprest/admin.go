package httprest

import (
	"net/http"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/http/respond"
)

type AdminHandler struct {
	respond *respond.Handle
}

func NewAdminHandler(respond *respond.Handle) *AdminHandler {
	return &AdminHandler{
		respond: respond,
	}
}

func (h *AdminHandler) Home(w http.ResponseWriter, r *http.Request) {
	h.respond.Success(w, map[string]string{"message": http.StatusText(http.StatusOK)}, http.StatusOK)
}
