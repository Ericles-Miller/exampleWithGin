package books

import (
	"exampleWithGin/internal/books/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type BookController struct{
	bookService models.BookService
}

func NewBookController(bookService models.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}


func (b *BookController) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")
	{
		books.GET("/", b.GetAllBooks)
		books.GET("/:id", b.GetBook)
		books.POST("", b.CreateBook)
		books.PUT("/:id", b.UpdateBook)
		books.DELETE("/:id", b.DeleteBook)
	}
}

func (b *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	createdBook, err := b.bookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	ctx.JSON(http.StatusCreated, createdBook)
}

func (b *BookController) GetBook(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	book, err := b.bookService.GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (b *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := b.bookService.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (b *BookController) UpdateBook(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	updatedBook, err := b.bookService.UpdateBook(id, &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
}

func (b *BookController) DeleteBook(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = b.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
