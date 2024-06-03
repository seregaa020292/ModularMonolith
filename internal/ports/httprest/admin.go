package httprest

import "net/http"

type AdminHandler struct{}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (handler *AdminHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
