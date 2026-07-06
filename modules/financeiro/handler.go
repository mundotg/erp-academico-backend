package financeiro

import (
	"net/http"
	"strconv"

	"erp-academico-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) *Handler { return &Handler{db: db} }

// ListarPagamentos godoc
// @Summary Lista pagamentos
// @Tags financeiro
// @Produce json
// @Success 200 {array} models.Pagamento
// @Router /financeiro/pagamentos [get]
func (h *Handler) ListarPagamentos(c *gin.Context) {
	var pagamentos []models.Pagamento
	h.db.Preload("Aluno").Find(&pagamentos)
	c.JSON(http.StatusOK, pagamentos)
}

// CriarPagamento godoc
// @Summary Cria pagamento
// @Tags financeiro
// @Accept json
// @Produce json
// @Param pagamento body models.Pagamento true "Pagamento"
// @Success 201 {object} models.Pagamento
// @Router /financeiro/pagamentos [post]
func (h *Handler) CriarPagamento(c *gin.Context) {
	var pagamento models.Pagamento
	if err := c.ShouldBindJSON(&pagamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	h.db.Create(&pagamento)
	c.JSON(http.StatusCreated, pagamento)
}

// ListarEncargos godoc
// @Summary Lista encargos financeiros
// @Tags financeiro
// @Produce json
// @Success 200 {array} models.Encargo
// @Router /financeiro/encargos [get]
func (h *Handler) ListarEncargos(c *gin.Context) {
	var encargos []models.Encargo
	h.db.Preload("Aluno").Find(&encargos)
	c.JSON(http.StatusOK, encargos)
}

// CriarEncargo godoc
// @Summary Cria encargo financeiro para aluno
// @Tags financeiro
// @Accept json
// @Produce json
// @Param encargo body models.Encargo true "Encargo"
// @Success 201 {object} models.Encargo
// @Failure 400 {object} map[string]string
// @Router /financeiro/encargos [post]
func (h *Handler) CriarEncargo(c *gin.Context) {
	var encargo models.Encargo
	if err := c.ShouldBindJSON(&encargo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if encargo.AlunoID == 0 || encargo.Tipo == "" || encargo.Descricao == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "aluno_id, tipo e descricao são obrigatórios"})
		return
	}
	if encargo.Valor <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "valor do encargo deve ser maior que zero"})
		return
	}
	if encargo.Status == "" {
		encargo.Status = "pendente"
	}

	h.db.Create(&encargo)
	c.JSON(http.StatusCreated, encargo)
}

// ValidarDividasAluno godoc
// @Summary Valida dívidas e encargos pendentes de um aluno
// @Tags financeiro
// @Produce json
// @Param alunoID path int true "ID do aluno"
// @Success 200 {object} models.ValidacaoDivida
// @Failure 400 {object} map[string]string
// @Router /financeiro/alunos/{alunoID}/dividas/validar [get]
func (h *Handler) ValidarDividasAluno(c *gin.Context) {
	alunoID, err := strconv.ParseUint(c.Param("alunoID"), 10, 64)
	if err != nil || alunoID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "alunoID inválido"})
		return
	}

	var pagamentos []models.Pagamento
	var encargos []models.Encargo
	h.db.Where("aluno_id = ? AND status = ?", alunoID, "pendente").Find(&pagamentos)
	h.db.Where("aluno_id = ? AND status = ?", alunoID, "pendente").Find(&encargos)

	resposta := montarValidacaoDivida(uint(alunoID), pagamentos, encargos)
	c.JSON(http.StatusOK, resposta)
}

func montarValidacaoDivida(alunoID uint, pagamentos []models.Pagamento, encargos []models.Encargo) models.ValidacaoDivida {
	var totalPagamentos float64
	for _, pagamento := range pagamentos {
		totalPagamentos += pagamento.Valor
	}

	var totalEncargos float64
	for _, encargo := range encargos {
		totalEncargos += encargo.Valor
	}

	total := totalPagamentos + totalEncargos
	return models.ValidacaoDivida{
		AlunoID:             alunoID,
		PossuiDivida:        total > 0,
		TotalPagamentos:     totalPagamentos,
		TotalEncargos:       totalEncargos,
		TotalEmDivida:       total,
		PagamentosPendentes: pagamentos,
		EncargosPendentes:   encargos,
	}
}
