package service

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/akshay-singla/employee-crud-api/employee/module"
	"github.com/akshay-singla/employee-crud-api/employee/repo"
	"github.com/gin-gonic/gin"
)

type Service struct {
	db *repo.Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: repo.NewRepository(db),
	}

}

func (s *Service) CreateEmployee(c *gin.Context) {
	var emp module.Employee
	if err := c.BindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.db.Create(emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (s *Service) UpdateEmployee(c *gin.Context) {
	var emp module.Employee
	if err := c.BindJSON(&emp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emp.ID = id
	err = s.db.Update(emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (s *Service) DeleteEmployee(c *gin.Context) {
	err := s.db.Delete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (s *Service) GetEmployeeByID(c *gin.Context) {
	emp, err := s.db.GetEmployeeByID(c.Param("id"))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emp)
}

func (s *Service) ListEmployee(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10 // Default page size
	}

	offset := (page - 1) * pageSize

	count, err := s.db.CountEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	list, err := s.db.ListEmployee(pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"page":          page,
			"page_size":     pageSize,
			"total_records": count,
			"list":          list,
		})
}
