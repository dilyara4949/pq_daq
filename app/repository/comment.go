package repository

import (
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
	// "gorm.io/gorm"
)

type CommentRepository interface {
	GetAll()([]*models.Comment, error)
	CreateComment(comment *models.Comment)(*models.Comment, error) 
	Delete(comment uint)( error)
	GetById(ID uint)(*models.Comment, error) 
	GetByProduct(product_id uint)([]*models.Comment, error)
}

type CommentRepo struct {
	Db *db.Database
}

func NewCommentRepo() CommentRepository {
	return &CommentRepo{Db: &db.DB}
}

func (s *CommentRepo)GetById(ID uint)(*models.Comment, error) {
	comment := models.Comment{}

	if result := s.Db.Db.Select("id", "comment").Where("id = ?", ID).First(&comment); result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}
func (s *CommentRepo)GetByProduct(product_id uint)([]*models.Comment, error) {
	comment := []*models.Comment{}

	if result := s.Db.Db.Where("Product_ID = ?", product_id).Find(&comment); result.Error != nil {
		return nil, result.Error
	}

	return comment, nil
}

func (s *CommentRepo)GetAll()([]*models.Comment, error) {
	var comments []*models.Comment

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

// func (s *CommentRepo)DeleteComment(comment *models.Comment)(*models.Comment, error) {
// 	if err := s.Db.Db.Delete(&comment).Error; err != nil {
// 		return comment, err
// 	}
// 	return comment, nil
// }

func (s *CommentRepo) Delete(id uint) error {
	if err := s.Db.Db.Delete(&models.Comment{}, id).Error; err != nil {
		return err
	}

	return nil
}


// package repository

// import (
// 	"github.com/dilyara4949/pq_daq/app/models"
// 	"github.com/dilyara4949/pq_daq/db"
// 	// "gorm.io/gorm"
// )

// type CommentRepository interface {
// 	GetAll() ([]*models.Comment, error)
// 	CreateComment(comment *models.Comment)(*models.Comment, error) 
// 	DeleteComment(comment *models.Comment) (*models.Comment, error)
// }

// type CommentRepo struct {
// 	Db *db.Database
// }

// func NewCommentRepo() CommentRepository {
// 	return &CommentRepo{Db: &db.DB}
// }

// func (s *CommentRepo)GetAll() ([]*models.Comment, error) {
// 	var comments []*models.Comment

// 	if result :=  s.Db.Db.Select("id", "UserID", "ProductID", "Comment").Find(&comments); result.Error != nil{
// 		return nil, result.Error
// 	}

// 	return comments, nil
// }

// func (s *CommentRepo)CreateComment(comment *models.Comment)(*models.Comment, error) {
// 	if err := s.Db.Db.Create(&comment).Error; err != nil {
// 		return comment, err
// 	}

// 	return comment, nil
// }

// func (s *CommentRepo)DeleteComment(comment *models.Comment) (*models.Comment, error) {
// 	if err := s.Db.Db.Delete(&comment).Error; err != nil {
// 		return comment, err
// 	}
// 	return comment, nil
// }