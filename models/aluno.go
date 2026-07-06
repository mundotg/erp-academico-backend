package models

import "gorm.io/gorm"

// Aluno agrupa os dados acadêmicos e pessoais do estudante.
type Aluno struct {
	gorm.Model
	Nome      string `json:"nome" gorm:"not null" example:"Ana Silva"`
	Email     string `json:"email" gorm:"uniqueIndex;not null" example:"ana.silva@example.com"`
	Matricula string `json:"matricula" gorm:"uniqueIndex;not null" example:"20260001"`
	CursoID   uint   `json:"curso_id" example:"1"`
	Curso     Curso  `json:"curso,omitempty"`
	Status    string `json:"status" gorm:"default:ativo" example:"ativo"`
}
