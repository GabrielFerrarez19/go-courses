package controllers

import "github.com/gin-gonic/gin"

type BooksController struct{}

func NewBooksController() *BooksController {
	return &BooksController{}
}

func (c *BooksController) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")

	{
		books.POST("", c.CreateBook)
		books.GET("", c.GetBooksByID)
		books.GET("/:id", c.ListBooks)
		books.PATCH("/:id", c.UpdateBook)
		books.DELETE("/:id", c.DeleteBook)
	}
}

func (c *BooksController) CreateBook(ctx *gin.Context) {}

func (c *BooksController) GetBooksByID(ctx *gin.Context) {}

func (c *BooksController) ListBooks(ctx *gin.Context) {}

func (c *BooksController) UpdateBook(ctx *gin.Context) {}

func (c *BooksController) DeleteBook(ctx *gin.Context) {}
