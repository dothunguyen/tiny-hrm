package routes

import (
	"tiny-hrm/src/handlers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Route define the enpoints and connect them to handler functions
func Route(router *gin.Engine, handlers *handlers.Handlers) {

	router.Use(static.Serve("/", static.LocalFile("./src/views/build", true)))

	employeeapi := router.Group("/api/v1/employees")
	{
		employeeapi.GET("/org", handlers.GetOrganisation)
		employeeapi.GET("/", handlers.GetAllEmployees)
		employeeapi.POST("/", handlers.AddEmployee)
		employeeapi.DELETE("/:id", handlers.DeleteEmployee)
	}
}
