package chatmodule

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// ChatService : Busines logic to fetch chats
type ChatService struct {
	db     *gorm.DB
	router *mux.Router
}

// NewChatService : Create a ChatService
func NewChatService(db *gorm.DB, router *mux.Router) *ChatService {
	return &ChatService{db, router}
}

// Get : Returns a Chat
func (c ChatService) Get() (*Chat, error) {
	chat := Chat{}
	err := c.db.First(chat).Error
	return &chat, err
}

// GetList : Returns a list of Chats
func (c ChatService) GetList() (*[]Chat, error) {
	chats := []Chat{}
	err := c.db.Find(&chats).Error

	return &chats, err
}
