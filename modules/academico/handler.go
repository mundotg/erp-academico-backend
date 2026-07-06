package academico

import (
	"net/http"

	"erp-academico-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) *Handler { return &Handler{db: db} }

// ListarAlunos godoc
// @Summary Lista alunos
// @Tags academico
// @Produce json
// @Success 200 {array} models.Aluno
// @Router /academico/alunos [get]
func (h *Handler) ListarAlunos(c *gin.Context) {
	var alunos []models.Aluno
	h.db.Preload("Curso").Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

// CriarAluno godoc
// @Summary Cria aluno
// @Tags academico
// @Accept json
// @Produce json
// @Param aluno body models.Aluno true "Aluno"
// @Success 201 {object} models.Aluno
// @Failure 400 {object} map[string]string
// @Router /academico/alunos [post]
func (h *Handler) CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	h.db.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)
}

// ListarCursos godoc
// @Summary Lista cursos
// @Tags academico
// @Produce json
// @Success 200 {array} models.Curso
// @Router /academico/cursos [get]
func (h *Handler) ListarCursos(c *gin.Context) {
	var cursos []models.Curso
	h.db.Preload("Disciplinas").Find(&cursos)
	c.JSON(http.StatusOK, cursos)
}

// CriarCurso godoc
// @Summary Cria curso
// @Tags academico
// @Accept json
// @Produce json
// @Param curso body models.Curso true "Curso"
// @Success 201 {object} models.Curso
// @Router /academico/cursos [post]
func (h *Handler) CriarCurso(c *gin.Context) {
	var curso models.Curso
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	h.db.Create(&curso)
	c.JSON(http.StatusCreated, curso)
}

// ListarNotas godoc
// @Summary Lista notas lançadas
// @Tags academico
// @Produce json
// @Success 200 {array} models.Nota
// @Router /academico/notas [get]
func (h *Handler) ListarNotas(c *gin.Context) {
	var notas []models.Nota
	h.db.Preload("Aluno").Preload("Disciplina").Find(&notas)
	c.JSON(http.StatusOK, notas)
}

// LancarNota godoc
// @Summary Lança nota de um aluno
// @Tags academico
// @Accept json
// @Produce json
// @Param nota body models.Nota true "Nota"
// @Success 201 {object} models.Nota
// @Failure 400 {object} map[string]string
// @Router /academico/notas [post]
func (h *Handler) LancarNota(c *gin.Context) {
	var nota models.Nota
	if err := c.ShouldBindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if nota.AlunoID == 0 || nota.DisciplinaID == 0 || nota.Avaliacao == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "aluno_id, disciplina_id e avaliacao são obrigatórios"})
		return
	}
	if nota.Valor < 0 || nota.Valor > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "valor da nota deve estar entre 0 e 20"})
		return
	}

	h.db.Create(&nota)
	c.JSON(http.StatusCreated, nota)
}
