package main

import (
	"log"

	booksController "librarymvc/internal/books/controllers"
	loansController "librarymvc/internal/loans/controllers"
	userController "librarymvc/internal/users/controllers"

	bookService "librarymvc/internal/books/services"
	loanService "librarymvc/internal/loans/services"
	userService "librarymvc/internal/users/services"

	bookRepository "librarymvc/internal/books/repository"
	loanRepository "librarymvc/internal/loans/repository"
	userRepository "librarymvc/internal/users/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Repositories
	loanRepo := loanRepository.NewLoansRepository()
	bookRepo := bookRepository.NewBookRepository()
	userRepo := userRepository.NewUserRepository()

	// Services
	bookSvc := bookService.NewBookService(bookRepo)
	userSvc := userService.NewUserService(userRepo)
	loanSvc := loanService.NewLoanService(loanRepo, bookSvc, userSvc)

	// Controllers
	booksController := booksController.NewBooksController(bookSvc)
	usersController := userController.NewUserController(userSvc)
	loansController := loansController.NewLoansController(loanSvc)

	booksController.RegisterRoutes(router)
	usersController.RegisterRoutes(router)
	loansController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
