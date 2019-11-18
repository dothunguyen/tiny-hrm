package main

import (
	"fmt"
	"tiny-hrm/handlers"
	"tiny-hrm/repository"
	"tiny-hrm/routes"
	"tiny-hrm/services"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	repo, err := repository.NewRepository()
	if err == nil {
		repo.InitDB()
		es := services.NewEmployeeService(repo)
		handlers := handlers.NewHandlers(es)

		router := gin.Default()
		routes.Route(router, handlers)

		router.Run(":3000")
	} else {
		fmt.Printf("Failed to connect to DB: %s", err)
	}

	defer repo.Close()
}
