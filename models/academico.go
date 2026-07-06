package models

import "gorm.io/gorm"

// Curso representa um curso ofertado pela instituição.
type Curso struct {
	gorm.Model
	Nome        string       `json:"nome" gorm:"not null" example:"Engenharia Informática"`
	Codigo      string       `json:"codigo" gorm:"uniqueIndex;not null" example:"EI"`
	Descricao   string       `json:"descricao" example:"Curso superior de engenharia informática"`
	Alunos      []Aluno      `json:"alunos,omitempty"`
	Disciplinas []Disciplina `json:"disciplinas,omitempty"`
}

// Disciplina representa uma unidade curricular do módulo acadêmico.
type Disciplina struct {
	gorm.Model
	Nome     string `json:"nome" gorm:"not null" example:"Programação I"`
	Codigo   string `json:"codigo" gorm:"uniqueIndex;not null" example:"PROG1"`
	Creditos uint   `json:"creditos" example:"6"`
	CursoID  uint   `json:"curso_id" example:"1"`
}

// Nota representa o lançamento de nota de um aluno em uma disciplina.
type Nota struct {
	gorm.Model
	AlunoID      uint       `json:"aluno_id" gorm:"not null;index" example:"1"`
	Aluno        Aluno      `json:"aluno,omitempty"`
	DisciplinaID uint       `json:"disciplina_id" gorm:"not null;index" example:"1"`
	Disciplina   Disciplina `json:"disciplina,omitempty"`
	Avaliacao    string     `json:"avaliacao" gorm:"not null" example:"Exame Final"`
	Valor        float64    `json:"valor" gorm:"not null" example:"16.5"`
	Observacao   string     `json:"observacao" example:"Aluno aprovado"`
	LancadoPor   string     `json:"lancado_por" example:"Professor Carlos"`
}
