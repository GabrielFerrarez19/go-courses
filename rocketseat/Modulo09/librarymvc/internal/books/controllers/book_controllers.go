package controllers

import (
	"net/http"
	"strconv"

	model "librarymvc/internal/books/models"

	"github.com/gin-gonic/gin"
)

type BooksController struct {
	bookService model.BookService
}

func NewBooksController(bookService model.BookService) *BooksController {
	return &BooksController{
		bookService: bookService,
	}
}

func (b *BooksController) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")

	{
		books.POST("", b.CreateBook)
		books.GET("", b.GetBooksByID)
		books.GET("/:id", b.ListBooks)
		books.PATCH("/:id", b.UpdateBook)
		books.DELETE("/:id", b.DeleteBook)
	}
}

func (b *BooksController) CreateBook(ctx *gin.Context) {
	var book model.Book

	if err := ctx.ShouldBind(book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := b.bookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (b *BooksController) GetBooksByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	book, err := b.bookService.GetBookByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (b *BooksController) ListBooks(ctx *gin.Context) {
	res, err := b.bookService.ListBook()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (b *BooksController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	var book model.Book

	if err := ctx.ShouldBind(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updateBook, err := b.bookService.UpdateBook(&book, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updateBook)
}

func (b *BooksController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	err = b.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
