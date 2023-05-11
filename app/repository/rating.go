package repository
import (
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
)

type RatingRepository interface {
	GetAll() ([]*models.Rating, error) 
	Create(rating *models.Rating) (*models.Rating, error)
	Delete(rating *models.Rating) (*models.Rating, error)
	GetById(ID uint) (*models.Rating, error) 
}

type RatingRepo struct {
	Db *db.Database
}

func NewRatingRepo() RatingRepository {
	return &RatingRepo{Db: &db.DB}
}




func (p *RatingRepo)GetAll() ([]*models.Rating, error) {
	var rating []*models.Rating

	if result := p.Db.Db.Find(&rating); result.Error != nil {
		return nil, result.Error
	}

	return rating, nil
}

func (p *RatingRepo)Create(rating *models.Rating) (*models.Rating, error) {

	if err := p.Db.Db.Create(&rating).Error; err != nil {
		return nil, err
	}

	return rating, nil
}

func (p *RatingRepo) Delete(rating *models.Rating) (*models.Rating, error) {
	if err := p.Db.Db.Delete(&rating).Error; err != nil {
		return rating, err
	}
	return rating, nil
}
func (c *RatingRepo) GetById(ID uint) (*models.Rating, error) {
	rating := models.Rating{}

	if result := c.Db.Db.Select("id", "name").Where("id = ?", ID).First(&rating); result.Error != nil {
		return nil, result.Error
	}

	return &rating, nil
}