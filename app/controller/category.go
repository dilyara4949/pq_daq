package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/app/services"
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)



type CategoryController struct {
	category service.CategoryService
}

func NewCategoryController(service service.CategoryService) *CategoryController {
	return &CategoryController{category: service}
}


func (p *CategoryController) GetAll(c *gin.Context) {
	categories, err := p.category.GetAll(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, categories)
}

func (p *CategoryController) GetById(c *gin.Context) {
	categoryId := c.Param("id")
	id, err:= strconv.Atoi(categoryId)

	if err != nil {
		log.Fatal("Id is incorrect")
	}

	categories, err := p.category.GetCategoryByID(c, uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, categories)
}

func (p *CategoryController) Create(c *gin.Context) {
	var category *models.Category
	if err := c.Bind(&category); err != nil {
		log.Fatal("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	category, err := p.category.CreateCategory(ctx, category)
	if err != nil {
		log.Fatal("Failed to create product", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, category)
}