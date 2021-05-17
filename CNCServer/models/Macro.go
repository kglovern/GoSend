package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Macro struct {
	gorm.Model
	Name      string    `json:"name"`
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

func getAllMacros(db *sql.DB) {
}

func getMacroById(db *sql.DB, id string) {}

func updateMacro(db *sql.DB, m *Macro) {}

func deleteMacro(db *sql.DB, id string) {}

func newMacro(db *sql.DB, m *Macro) {}
