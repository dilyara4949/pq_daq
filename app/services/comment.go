package service

import (
	"context"

	"github.com/dilyara4949/pq_daq/app/models"
	repo "github.com/dilyara4949/pq_daq/app/repository"
)

type commentService struct {
	comment repo.CommentRepo
}

func (c *commentService)GetAllComments(ctx context.Context, comment *models.Comment) ([]models.Comment, error) {
	comments, err := c.comment.GetAllComments()
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