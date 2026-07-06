package models

import (
	"time"

	"gorm.io/gorm"
)

// Pagamento representa uma cobrança ou pagamento associado ao aluno.
type Pagamento struct {
	gorm.Model
	AlunoID        uint      `json:"aluno_id" example:"1"`
	Aluno          Aluno     `json:"aluno,omitempty"`
	Descricao      string    `json:"descricao" gorm:"not null" example:"Mensalidade Julho"`
	Valor          float64   `json:"valor" gorm:"not null" example:"25000"`
	DataVencimento time.Time `json:"data_vencimento" example:"2026-07-31T00:00:00Z"`
	Status         string    `json:"status" gorm:"default:pendente" example:"pendente"`
}

// Encargo representa uma taxa, multa ou outro encargo financeiro lançado ao aluno.
type Encargo struct {
	gorm.Model
	AlunoID        uint      `json:"aluno_id" gorm:"not null;index" example:"1"`
	Aluno          Aluno     `json:"aluno,omitempty"`
	Tipo           string    `json:"tipo" gorm:"not null" example:"multa"`
	Descricao      string    `json:"descricao" gorm:"not null" example:"Multa por atraso na mensalidade"`
	Valor          float64   `json:"valor" gorm:"not null" example:"1500"`
	DataVencimento time.Time `json:"data_vencimento" example:"2026-07-31T00:00:00Z"`
	Status         string    `json:"status" gorm:"default:pendente" example:"pendente"`
}

// ValidacaoDivida resume a situação financeira pendente de um aluno.
type ValidacaoDivida struct {
	AlunoID             uint        `json:"aluno_id" example:"1"`
	PossuiDivida        bool        `json:"possui_divida" example:"true"`
	TotalPagamentos     float64     `json:"total_pagamentos" example:"25000"`
	TotalEncargos       float64     `json:"total_encargos" example:"1500"`
	TotalEmDivida       float64     `json:"total_em_divida" example:"26500"`
	PagamentosPendentes []Pagamento `json:"pagamentos_pendentes"`
	EncargosPendentes   []Encargo   `json:"encargos_pendentes"`
}
