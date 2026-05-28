package books

import (
	"exampleWithGin/internal/books/models"
	"exampleWithGin/pkg/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController struct {
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
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.Book]("Invalid request body"))
		return
	}

	createdBook, err := b.bookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httputil.Fail[*models.Book]("Failed to create book"))
		return
	}

	ctx.JSON(http.StatusCreated, httputil.Success(createdBook))
}

func (b *BookController) GetBook(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.Book]("Invalid UUID"))
		return
	}

	book, err := b.bookService.GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, httputil.Fail[*models.Book]("Book not found"))
		return
	}

	ctx.JSON(http.StatusOK, httputil.Success(book))
}

func (b *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := b.bookService.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httputil.Fail[[]*models.Book]("Failed to retrieve books"))
		return
	}

	ctx.JSON(http.StatusOK, httputil.Success(books))
}

func (b *BookController) UpdateBook(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.Book]("Invalid UUID"))
		return
	}

	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.Book]("Invalid request body"))
		return
	}

	updatedBook, err := b.bookService.UpdateBook(id, &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httputil.Fail[*models.Book]("Failed to update book"))
		return
	}

	ctx.JSON(http.StatusOK, httputil.Success(updatedBook))
}

func (b *BookController) DeleteBook(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.Book]("Invalid UUID"))
		return
	}

	err = b.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, httputil.Fail[*models.Book]("Book not found"))
		return
	}

	ctx.Status(http.StatusNoContent)
}
