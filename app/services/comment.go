package service

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)
type CommentService interface {
	GetAll(ctx context.Context, comment *models.Comment) ([]*models.Comment, error)
	DeleteComment(ctx context.Context, comment *models.Comment)(*models.Comment, error)
	CreateComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) 
}

type commentService struct {
	comment repo.CommentRepository
}
func NewCommentService(repo repo.CommentRepository) CommentService {
	return &commentService{comment: repo}
}

func (c *commentService)GetAll(ctx context.Context, comment *models.Comment) ([]*models.Comment, error) {
	comments, err := c.comment.GetAll()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *commentService) DeleteComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) {
	comment, err := c.comment.DeleteComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentService) CreateComment(ctx context.Context, comment *models.Comment)(*models.Comment, error) {
	comment, err := c.comment.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
