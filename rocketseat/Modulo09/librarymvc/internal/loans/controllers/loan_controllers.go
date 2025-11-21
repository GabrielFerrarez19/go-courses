package controllers

import (
	"net/http"
	"strconv"

	"librarymvc/internal/loans/models"

	"github.com/gin-gonic/gin"
)

type LoansController struct {
	loanService models.LoansService
}

func NewLoansController(loanService models.LoansService) *LoansController {
	return &LoansController{
		loanService: loanService,
	}
}

func (l *LoansController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")

	{
		loans.POST("", l.CreateLoans)
		loans.GET("", l.GetLoansByID)
		loans.GET("/:id", l.ListLoans)
		loans.GET("/user/:userId", l.ListLoansByUserID)
		loans.GET("/return/:loanId", l.ReturnBook)
	}
}

func (l *LoansController) CreateLoans(ctx *gin.Context) {
	var req struct {
		BookID int64 `json:"bookID"`
		UserID int64 `json:"userID"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := l.loanService.CreateLoan(req.BookID, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (l *LoansController) GetLoansByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	book, err := l.loanService.GetLoanByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (l *LoansController) ListLoans(ctx *gin.Context) {
	res, err := l.loanService.ListLoan()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (l *LoansController) ListLoansByUserID(ctx *gin.Context) {
	userID, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	res, err := l.loanService.ListActiveLoanByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (l *LoansController) ReturnBook(ctx *gin.Context) {
	loanID, err := strconv.ParseInt(ctx.Param("loanId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	err = l.loanService.ReturnBook(loanID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
