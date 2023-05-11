package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/app/services"
	"github.com/gin-gonic/gin"
)



type RatingController struct {
	rating service.RatingService
}

func NewRatingController(service service.RatingService) *RatingController {
	return &RatingController{rating: service}
}


func (p *RatingController) GetAll(c *gin.Context) {
	ratings, err := p.rating.GetAll(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, ratings)
}

func (p *RatingController) GetById(c *gin.Context) {
	id := c.Param("product_id")
	productId, err:= strconv.Atoi(id)

	if err != nil {
		fmt.Println(id)
		log.Fatal("Id is incorrect")
	}
	
	rating, err := p.rating.GetRatingByID(c, uint(productId))
	view := models.RatingRep{
		ProductId: rating.ProductID,
	}
	if rating.Amount == 0{
		view.Rating = 0
	} else {
		view.Rating = float64(rating.Sum) / float64(rating.Amount)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, view)
}

func (p *RatingController) Create(c *gin.Context) {
	var rating *models.RatingRep
	if err := c.Bind(&rating); err != nil {
		log.Fatal("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	r, err := p.rating.GetRatingByID(ctx, rating.ProductId)

	if err != nil {
		log.Fatal(err)
	}

	r.Amount = r.Amount + 1
	r.Sum += int(rating.Rating)
	r, err = p.rating.Update(ctx, r)
	if err != nil {
		log.Fatal("Failed to create product", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	product, err := p.rating.GetProductByID(ctx, uint(rating.ProductId))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	product.Rating = float64(r.Sum) / float64(r.Amount)

	product, err = p.rating.UpdateProduct(ctx, product)
	c.JSON(http.StatusOK, product)
}