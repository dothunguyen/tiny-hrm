package handlers

import (
	"net/http"
	"tiny-hrm/src/repository"
	"tiny-hrm/src/services"

	"github.com/gin-gonic/gin"
)

// Handlers provide router with handler functions
type Handlers struct {
	es *services.EmployeeService
}

// NewHandlers create new instance of Services
func NewHandlers(es *services.EmployeeService) *Handlers {
	return &Handlers{es: es}
}

// GetOrganisation GET org handler
func (s *Handlers) GetOrganisation(c *gin.Context) {
	org, err := s.es.GetOrganisation()
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, org)
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

// GetAllEmployees GET all employees handler
func (s *Handlers) GetAllEmployees(c *gin.Context) {
	employees, err := s.es.GetAllEmployees()
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, employees)
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(
			http.StatusNoContent,
			gin.H{})
	}
}

// AddEmployee POST employee Handler
func (s *Handlers) AddEmployee(c *gin.Context) {
	var e repository.Employee
	bindingerr := c.BindJSON(&e)
	if bindingerr == nil {
		rs, err := s.es.AddEmployee(e)
		if err == nil {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusCreated, rs)
		} else {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotModified, gin.H{"message": "Couldnot insert employee record"})
		}
	}
}

// DeleteEmployee POST employee Handler
func (s *Handlers) DeleteEmployee(c *gin.Context) {
	err := s.es.DeleteEmployee(c.Param("id"))
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"message": "record has been deleted"})
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusNotModified, gin.H{"message": "Could not delete employee record"})
	}
}
