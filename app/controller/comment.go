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

type CommentController struct {
	comment service.CommentService
}

func NewCommentController(service service.CommentService) *CommentController {
	return &CommentController{comment: service}
}

func (p *CommentController) GetById(c *gin.Context) {
	commentID := c.Param("id")
	id, err := strconv.Atoi(commentID)

	if err != nil {
		log.Fatal("Id is incorrect")
	}
	comments, err := p.comment.GetByID(c, uint(id))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, comments)
}


func (p *CommentController) AddComment(c *gin.Context) {
	var comment *models.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	createdComment, err := p.comment.CreateComment(ctx, comment)
	if err != nil {
		log.Fatal("Failed to create comment: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, createdComment)
}



func (p *CommentController) DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	id, err := strconv.Atoi(commentID)

	if err != nil {
		log.Fatal("Id is incorrect")
	}

	err = p.comment.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, "Comment deleted")
}


func (p *CommentController) GetAll(c *gin.Context) {
	comment, err := p.comment.GetAll(c, )

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, comment)
}


func (p *CommentController) GetByProduct(c *gin.Context) {
	product_id := c.Param("product_id")
	id, err := strconv.Atoi(product_id)

	if err != nil {
		log.Fatal("Id is incorrect")
	}
	comments, err := p.comment.GetByProduct(c, uint(id))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, comments)
}