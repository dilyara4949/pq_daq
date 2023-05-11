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
	UpdatRating(rating *models.Rating) (*models.Rating, error)
	UpdatProduct(product *models.Product) (*models.Product, error)
	GetByIdProduct(ID uint) (*models.Product, error)
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
func (c *RatingRepo) GetById(productId uint) (*models.Rating, error) {
	rating := models.Rating{}

	if result := c.Db.Db.Where("product_ID = ?", productId).First(&rating); result.Error != nil {
		return nil, result.Error
	}

	return &rating, nil
}

func (s *RatingRepo)UpdatRating(rating *models.Rating) (*models.Rating, error) {
	if err := s.Db.Db.Save(&rating).Error; err != nil {
		return rating, err
	}

	return rating, nil	
}

func (p *RatingRepo) GetByIdProduct(ID uint) (*models.Product, error) {
	product := models.Product{}

	if result := p.Db.Db.Where("id = ?", ID).First(&product); result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}
func (s *RatingRepo)UpdatProduct(product *models.Product) (*models.Product, error) {
	if err := s.Db.Db.Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil	
}