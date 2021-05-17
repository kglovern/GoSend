package models

import (
	_ "GoSend/CNCServer"
	"github.com/jinzhu/gorm"
)

type Macro struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	CreatedAt  string `json:"created_at"`
	Content    string `json:"content"`
}

func getAllMacros(db *gorm.DB) {
}

func getMacroById(db *gorm.DB, id string) {}

func updateMacro(db *gorm.DB, m *Macro) {}

func deleteMacro(db *gorm.DB, id string) {}

func newMacro(db *gorm.DB, m *Macro) {}
