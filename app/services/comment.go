package service

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)
type CommentService interface {
	GetAll(ctx context.Context) ([]*models.Comment, error)
	// DeleteComment(ctx context.Context, comment *models.Comment)(*models.Comment, error)
	CreateComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) 
	GetByID(ctx context.Context, id uint)(*models.Comment, error)
	Delete(id int) error
	GetByProduct(ctx context.Context, product_id uint)([]*models.Comment, error)
}

type commentService struct {
	comment repo.CommentRepository
}
func NewCommentService(repo repo.CommentRepository) CommentService {
	return &commentService{comment: repo}
}

func (c *commentService) GetByID(ctx context.Context, id uint)(*models.Comment, error) {
	comment, err := c.comment.GetById(id)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
func (c *commentService) GetByProduct(ctx context.Context, product_id uint)([]*models.Comment, error) {
	comments, err := c.comment.GetByProduct(product_id)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *commentService)GetAll(ctx context.Context) ([]*models.Comment, error) {
	comments, err := c.comment.GetAll()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// func (c *commentService) DeleteComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) {
// 	comment, err := c.comment.DeleteComment(comment)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return comment, nil
// }

func (c *commentService) Delete(id int) error {
	err := c.comment.Delete(uint(id))
	return err
}


func (c *commentService) CreateComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) {
	comment, err := c.comment.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}


// package service

// import (
// 	"context"

// 	"github.com/dilyara4949/pq_daq/app/models"
// 	repo "github.com/dilyara4949/pq_daq/app/repository"
// )
// type CommentService interface {
// 	GetAll(ctx context.Context, comment *models.Comment) ([]*models.Comment, error)
// 	DeleteComment(ctx context.Context, comment *models.Comment)(*models.Comment, error)
// 	CreateComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) 
// }

// type commentService struct {
// 	comment repo.CommentRepository
// }
// func NewCommentService(repo repo.CommentRepository) CommentService {
// 	return &commentService{comment: repo}
// }

// func (c *commentService)GetAll(ctx context.Context, comment *models.Comment) ([]*models.Comment, error) {
// 	comments, err := c.comment.GetAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return comments, nil
// }

// func (c *commentService) DeleteComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) {
// 	comment, err := c.comment.DeleteComment(comment)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return comment, nil
// }

// func (c *commentService) CreateComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) {
// 	comment, err := c.comment.CreateComment(comment)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return comment, nil
// }
