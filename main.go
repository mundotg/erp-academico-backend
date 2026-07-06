package main

import (
	"log"

	"erp-academico-backend/config"
	_ "erp-academico-backend/docs"
	"erp-academico-backend/routes"
)

// @title ERP Acadêmico API
// @version 1.0
// @description API REST para gestão acadêmica, recursos humanos e financeiro.
// @termsOfService http://swagger.io/terms/
// @contact.name Suporte API
// @contact.email suporte@example.com
// @license.name MIT
// @BasePath /api/v1
func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("erro ao conectar ao banco de dados: %v", err)
	}

	router := routes.SetupRouter(db)
	if err := router.Run(); err != nil {
		log.Fatalf("erro ao iniciar servidor: %v", err)
	}
}
