package controllers

import "github.com/gin-gonic/gin"

type LoansController struct{}

func NewLoansController() *LoansController {
	return &LoansController{}
}

func (c *LoansController) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/loans")

	{
		books.POST("", c.CreateLoans)
		books.GET("", c.GetLoansByID)
		books.GET("/:id", c.ListLoans)
		books.PATCH("/:id", c.UpdateLoans)
		books.DELETE("/:id", c.DeleteLoans)
	}
}

func (c *LoansController) CreateLoans(ctx *gin.Context) {}

func (c *LoansController) GetLoansByID(ctx *gin.Context) {}

func (c *LoansController) ListLoans(ctx *gin.Context) {}

func (c *LoansController) UpdateLoans(ctx *gin.Context) {}

func (c *LoansController) DeleteLoans(ctx *gin.Context) {}
