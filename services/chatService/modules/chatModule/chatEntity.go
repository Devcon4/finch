package chatmodule

import "github.com/jinzhu/gorm"

// Chat Entity
type Chat struct {
	gorm.Model
	Message string `json:"message"`
}
