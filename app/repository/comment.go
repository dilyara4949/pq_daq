package repository

import (
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
	// "gorm.io/gorm"
)

type CommentRepo struct {
	Db *db.Database
}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{Db: &db.DB}
}

func (s *CommentRepo)GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment

	if result :=  s.Db.Db.Select("id", "UserID", "ProductID", "Comment").Find(&comments); result.Error != nil{
		return nil, result.Error
	}

	return comments, nil
}

func (s *CommentRepo)CreateComment(comment *models.Comment)(*models.Comment, error) {
	if err := s.Db.Db.Create(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}

func (s *CommentRepo)DeleteComment(comment *models.Comment) (*models.Comment, error) {
	if err := s.Db.Db.Delete(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}