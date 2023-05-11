package controller

import (
	//  "log"
	"net/http"
	"strconv"
	//  "strconv"

	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/app/services"
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

type OrderController struct {
 order service.OrderService
}

type Payment struct {
	Money string
}
func NewOrderController(service service.OrderService) *OrderController {
 return &OrderController{order: service}
}

func (o *OrderController) GetAllOrders(c *gin.Context) {
	orders, err := o.order.GetAll(c, &models.Order{})
	
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, orders)
	
}

func (o *OrderController) GetByUser(c *gin.Context) {
	id := c.Param("user_id")
	user_id, err:= strconv.Atoi(id)

	if err != nil {
		panic("Id is incorrect")
	}

	order, err := o.order.GetByUser(c, uint(user_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, order)

}

func (o *OrderController) GetByProduct(c *gin.Context) {
	id := c.Param("product_id")
	product_id, err:= strconv.Atoi(id)

	if err != nil {
		panic("Id is incorrect")
	}

	order, err := o.order.GetByProduct(c, uint(product_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, order)
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var order *models.Order
	if err := c.Bind(&order); err != nil {
		// log.Fatal("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	product, err := o.order.GetProductByID(ctx, uint(order.ProductId))
	order.TotalPrice = product.Price
	
	order, err = o.order.CreateOrder(ctx, order)
	if err != nil {
		// log.Fatal("Failed to create product", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)

}
func (o *OrderController) Pay(c *gin.Context) {
	var money Payment
	
	if err := c.Bind(&money); err != nil {
		// log.Fatal("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	order_id, err:= strconv.Atoi(id)

	if err != nil {
		panic("Id is incorrect")
	}
	ctx := c.Request.Context()
	if money.Money == "pay" {
		order, err := o.order.GetOrderByID(ctx, uint(order_id))		
		order.PaymentStatus = "paid"
		order, err = o.order.Update(ctx, order)
		if err != nil {
			// log.Fatal("Failed to create product", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, order)
	} else {
		c.JSON(http.StatusOK, "Pay your order")
	}
	
	

	
}
func (o *OrderController) Delete(c *gin.Context) {

}