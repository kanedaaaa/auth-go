package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kanedaaaa/auth-go/types"
	"github.com/kanedaaaa/auth-go/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin)
	router.HandleFunc("/signup", h.handleSignup)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {
	var payload types.SignupPyaload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

}
