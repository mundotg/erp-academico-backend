package models

import "gorm.io/gorm"

// Departamento representa uma área administrativa da instituição.
type Departamento struct {
	gorm.Model
	Nome         string        `json:"nome" gorm:"uniqueIndex;not null" example:"Secretaria Acadêmica"`
	Descricao    string        `json:"descricao" example:"Departamento responsável por atendimento acadêmico"`
	Funcionarios []Funcionario `json:"funcionarios,omitempty"`
}

// Funcionario representa colaboradores docentes e administrativos.
type Funcionario struct {
	gorm.Model
	Nome           string       `json:"nome" gorm:"not null" example:"Carlos Mendes"`
	Email          string       `json:"email" gorm:"uniqueIndex;not null" example:"carlos.mendes@example.com"`
	Cargo          string       `json:"cargo" example:"Professor"`
	DepartamentoID uint         `json:"departamento_id" example:"1"`
	Departamento   Departamento `json:"departamento,omitempty"`
	Ativo          bool         `json:"ativo" gorm:"default:true" example:"true"`
}
