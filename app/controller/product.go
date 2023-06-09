package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/app/services"
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	// "github.com/jinzhu/copier"
)

type sort struct {
	Sort string `json:"sort"`
}


type ProductController struct {
	product service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{product: service}
}


func (p *ProductController) GetAll(c *gin.Context) {
	products, err := p.product.GetAll(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, products)
}
func (p *ProductController) GetById(c *gin.Context) {
	productId := c.Param("id")
	id, err:= strconv.Atoi(productId)
	
	if err != nil {
		panic("Id is incorrect")
	}
	fmt.Println(productId)
	product, err := p.product.GetProductByID(c, uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, product)
}


func (p *ProductController) getCategortById(c *gin.Context, id uint)( *models.Category, error){
	return p.product.GetCategorytByID(c, id)
}

// func (p *ProductController) CreateProduct(c *gin.Context) {
// 	var item models.Product
// 	if err := c.Bind(&item); err != nil {
// 		log.Fatal("Failed to parse request body: ", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// validate := validator.New()
// 	// err := validate.Struct(item)
// 	// if err != nil {
// 	// 	log.Fatal("Request body is invalid: ", err.Error())
// 	// 	c.JSON(http.StatusBadRequest, err.Error())
// 	// 	return
// 	// }

// 	ctx := c.Request.Context()
// 	products, err := p.product.CreateProduct(ctx, &item)
// 	if err != nil {
// 		log.Fatal("Failed to create product", err.Error())
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	// var res models.Product
// 	// copier.Copy(&res, &products)
// 	c.JSON(http.StatusOK, products)
// }

// func (p *ProductController) crRAT(*models.Rating) {
	
// }

func (p *ProductController) CreateProduct(c *gin.Context) {
	var item models.Product
	err := c.Bind(&item)
	if err != nil {
	    log.Fatal("Failed to parse request body: ", err)
	    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
	}
	category, err := p.getCategortById(c, uint(item.CategoryID))
	item.Category = *category
	ctx := c.Request.Context()
	product, err := p.product.CreateProduct(ctx, &item)
	if err != nil {
	    log.Fatal("Failed to create product: ", err.Error())
	    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
	}

	rating := models.Rating{}
	rating.Product = item
	rating.ProductID = item.ID
	rating.Amount = 0
	rating.Sum = 0
	p.product.CreateRating(ctx, &rating)
    
	c.JSON(http.StatusOK, product)
    }
    
func (p *ProductController) SortByPrice(c *gin.Context) {
	sort := sort{}
	ctx := c.Request.Context()
	// err := c.Bind(sort)
	if err := c.ShouldBindJSON(&sort);err != nil {
	    log.Fatal("Failed to parse request body: ", err)
	    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
	}
	fmt.Println(sort)
	products, err := p.product.SortByPrice(ctx, sort.Sort)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, products)
}

func (p *ProductController) SortByRating(c *gin.Context) {
	sort := sort{}
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&sort);err != nil {
	    log.Fatal("Failed to parse request body: ", err)
	    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
	}
	
	fmt.Println(sort)
	products, err := p.product.SortByRating(ctx, sort.Sort)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, products)
}