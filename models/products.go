package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        int `json:"id" gorm:"primarykey"`
	Nome       string `json:"nome"`
	Descricao  string `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int `json:"quantidade"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}