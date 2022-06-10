package model

import (
	"time"
)

type Produto struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	Codigo    string    `json:"codigo" gorm:"uniqueIndex"`
	Nome      string    `json:"nome"`
	EstoqueID uint      `json:"-"`
	Estoque   Estoque   `json:"estoque" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PrecoDe   float64   `json:"preco_de"`
	PrecoPor  float64   `json:"preco_por"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Estoque struct {
	ID                uint      `json:"-" gorm:"primaryKey"`
	EstoqueTotal      int64     `json:"estoque_total"`
	EstoqueCorte      int64     `json:"estoque_corte"`
	EstoqueDisponivel int64     `json:"estoque_disponivel"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (p Produto) PrecoInvalido() bool {
	return p.PrecoPor > p.PrecoDe
}

func (p Produto) EstoqueInvalido() bool {
	return p.Estoque.EstoqueCorte > p.Estoque.EstoqueTotal
}
