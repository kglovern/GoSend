package models

import (
	"database/sql"
	"time"
)

type Macro struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

func getAllMacros(db *sql.DB) {
}

func getMacroById(db *sql.DB, id string) {}

func updateMacro(db *sql.DB, m *Macro) {}

func deleteMacro(db *sql.DB, id string) {}

func newMacro(db *sql.DB, m *Macro) {}
