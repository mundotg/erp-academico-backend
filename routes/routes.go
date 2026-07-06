package routes

import (
	"erp-academico-backend/modules/academico"
	"erp-academico-backend/modules/financeiro"
	"erp-academico-backend/modules/rh"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// SetupRouter registra os módulos da API e a documentação Swagger.
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	registrarAcademico(api, db)
	registrarRH(api, db)
	registrarFinanceiro(api, db)

	return router
}

func registrarAcademico(api *gin.RouterGroup, db *gorm.DB) {
	h := academico.NewHandler(db)
	academicoGroup := api.Group("/academico")
	academicoGroup.GET("/alunos", h.ListarAlunos)
	academicoGroup.POST("/alunos", h.CriarAluno)
	academicoGroup.GET("/cursos", h.ListarCursos)
	academicoGroup.POST("/cursos", h.CriarCurso)
	academicoGroup.GET("/notas", h.ListarNotas)
	academicoGroup.POST("/notas", h.LancarNota)
}

func registrarRH(api *gin.RouterGroup, db *gorm.DB) {
	h := rh.NewHandler(db)
	rhGroup := api.Group("/rh")
	rhGroup.GET("/funcionarios", h.ListarFuncionarios)
	rhGroup.POST("/funcionarios", h.CriarFuncionario)
}

func registrarFinanceiro(api *gin.RouterGroup, db *gorm.DB) {
	h := financeiro.NewHandler(db)
	financeiroGroup := api.Group("/financeiro")
	financeiroGroup.GET("/pagamentos", h.ListarPagamentos)
	financeiroGroup.POST("/pagamentos", h.CriarPagamento)
	financeiroGroup.GET("/encargos", h.ListarEncargos)
	financeiroGroup.POST("/encargos", h.CriarEncargo)
	financeiroGroup.GET("/alunos/:alunoID/dividas/validar", h.ValidarDividasAluno)
}
