module erp-academico-backend

go 1.22

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.4
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.25.12
)

replace github.com/gin-gonic/gin => ./internal/stubs/gin

replace github.com/swaggo/files => ./internal/stubs/swaggerfiles

replace github.com/swaggo/gin-swagger => ./internal/stubs/ginswagger

replace github.com/swaggo/swag => ./internal/stubs/swag

replace gorm.io/driver/sqlite => ./internal/stubs/sqlite

replace gorm.io/gorm => ./internal/stubs/gorm
