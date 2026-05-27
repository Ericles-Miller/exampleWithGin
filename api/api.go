package api

import (
	"exampleWithGin/internal/books"
	bookService "exampleWithGin/internal/books/service"
	"exampleWithGin/internal/loans"
	loanService "exampleWithGin/internal/loans/services"
	"exampleWithGin/internal/users"
	userService "exampleWithGin/internal/users/services"
	"exampleWithGin/internal/users/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewServer(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	// Users
	uRepo := repositories.NewUserRepository(pool)
	uService := userService.NewUserService(uRepo)
	uController := users.NewUserController(uService)
	uController.RegisterRoutes(router)

	// Books
	// bRepo := bookRepositories.NewBookRepository(pool)
	bService := bookService.NewBookService(nil) // substituir nil pelo repo quando implementado
	bController := books.NewBookController(bService)
	bController.RegisterRoutes(router)

	// Loans
	// lRepo := loanRepositories.NewLoanRepository(pool)
	lService := loanService.NewLoanService(nil, bService, uService) // substituir nil pelo repo quando implementado
	lController := loans.NewLoanController(lService)
	lController.RegisterRoutes(router)

	return router
}
