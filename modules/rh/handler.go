package rh

import (
	"net/http"

	"erp-academico-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) *Handler { return &Handler{db: db} }

// ListarFuncionarios godoc
// @Summary Lista funcionários
// @Tags rh
// @Produce json
// @Success 200 {array} models.Funcionario
// @Router /rh/funcionarios [get]
func (h *Handler) ListarFuncionarios(c *gin.Context) {
	var funcionarios []models.Funcionario
	h.db.Preload("Departamento").Find(&funcionarios)
	c.JSON(http.StatusOK, funcionarios)
}

// CriarFuncionario godoc
// @Summary Cria funcionário
// @Tags rh
// @Accept json
// @Produce json
// @Param funcionario body models.Funcionario true "Funcionário"
// @Success 201 {object} models.Funcionario
// @Router /rh/funcionarios [post]
func (h *Handler) CriarFuncionario(c *gin.Context) {
	var funcionario models.Funcionario
	if err := c.ShouldBindJSON(&funcionario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	h.db.Create(&funcionario)
	c.JSON(http.StatusCreated, funcionario)
}
