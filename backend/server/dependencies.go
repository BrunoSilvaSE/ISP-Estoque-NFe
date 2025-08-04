package server

import (
	"backend/auth"
	"backend/database"
	"backend/handlers"
	"backend/middleware"
	"backend/repositories"
	"backend/services"
	"database/sql"
	"log"
	"os"
)

type Dependencies struct {
	Middleware		*middleware.AuthMiddleware

	UserHandler     *handlers.UserHandler
	TypeHandler     *handlers.TypeHandler

	AuthHandler		*auth.AuthHandler
	 
}

func BuildDependencies(db *sql.DB) *Dependencies {
	chaveJwt := os.Getenv("JWT_KEY_INDESCOBRIVEL_INDECIFRAVEL_INDESCOBERTA")
	if chaveJwt == "" {
		chaveJwt = "JWT_KEY_INDESCOBRIVEL_INDECIFRAVEL_INDESCOBERTA_2"
		log.Println("Chave JWT não encontrada, usando chave padrão")
	}

	dbCliente := &database.DatabaseCliente{DB: db}

 	userRepo := repositories.NewUserRepository(dbCliente)
 	userService := services.NewUserService(userRepo)
 	userHandler := handlers.NewUserHandler(userService)

	typeRep := repositories.NewTypeRepository(dbCliente)
	typeService := services.NewTypeService(typeRep)
	typeHandler := handlers.NewTypeHandler(typeService)

	authService := auth.NewAutenticacaoService(userRepo, []byte(chaveJwt))
	authHandler := auth.NewAuthHandler(authService)

	middleware := middleware.NewMiddlewareAuth(authService)

	return &Dependencies{
		Middleware: middleware,

		UserHandler: userHandler,
		TypeHandler: typeHandler,

		AuthHandler: authHandler,
	}
}