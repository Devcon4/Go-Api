package chatmodule

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// ChatService : Busines logic to fetch chats
type ChatService struct {
	db     *sqlx.DB
	router *mux.Router
}

// NewChatService : Create a ChatService
func NewChatService(db *sqlx.DB, router *mux.Router) *ChatService {
	return &ChatService{db, router}
}

// Get : Returns a Chat
func (c ChatService) Get() (*Chat, error) {

	return &Chat{
		2,
		"Test again!",
	}, nil
}

// GetList : Returns a list of Chats
func (c ChatService) GetList() (*[]Chat, error) {
	chats := []Chat{}
	err := c.db.Select(&chats, "select * from chat")

	return &chats, err
}
