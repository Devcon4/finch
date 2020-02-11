package chatmodule

import (
	"net/http"

	"github.com/Devcon4/finch/services/chatService/framework"
	"github.com/gorilla/mux"
)

// ChatHandler : Controller for chats
type ChatHandler struct {
	router      *mux.Router
	chatService *ChatService
}

// NewChatHandler : Creates a ChatHandler
func NewChatHandler(router *mux.Router, chatService *ChatService) *ChatHandler {
	return &ChatHandler{router, chatService}
}

// Register : Wires up route handlers
func (h ChatHandler) Register() {
	h.router.HandleFunc("/chat", h.Get).Methods("GET")
	h.router.HandleFunc("/chats", h.GetList).Methods("GET")
}

// Get : Get handler func
func (h ChatHandler) Get(w http.ResponseWriter, r *http.Request) {
	c, err := h.chatService.Get()

	framework.JSONHandler(w, c, err)
}

// GetList : GetList handler func
func (h ChatHandler) GetList(w http.ResponseWriter, r *http.Request) {
	l, err := h.chatService.GetList()

	framework.JSONHandler(w, l, err)
}
