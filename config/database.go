package config

import (
	"os"

	"erp-academico-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDatabase abre a conexão com o banco de dados e executa as migrações.
func ConnectDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "erp_academico.db"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.Aluno{},
		&models.Curso{},
		&models.Disciplina{},
		&models.Nota{},
		&models.Departamento{},
		&models.Funcionario{},
		&models.Pagamento{},
		&models.Encargo{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
