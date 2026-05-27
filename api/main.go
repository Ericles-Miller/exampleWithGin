package main

import (
	"exampleWithGin/internal/books"
	"exampleWithGin/internal/loans"
	"exampleWithGin/internal/users"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	usersController := users.NewUserController()
	usersController.RegisterRoutes(router)

	booksController := books.NewBookController()
	booksController.RegisterRoutes(router)

	loansController := loans.NewLoanController()
	loansController.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}