package main

import (
	"log"

	bookscontroller "librarymvc/internal/books/controllers"
	loanscontroller "librarymvc/internal/loans/controllers"
	usercontroller "librarymvc/internal/users/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	booksController := bookscontroller.NewBooksController()
	usersController := usercontroller.NewUserController()
	loansController := loanscontroller.NewLoansController()

	booksController.RegisterRoutes(router)
	usersController.RegisterRoutes(router)
	loansController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
