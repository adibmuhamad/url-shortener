package shorter

import "gorm.io/gorm"

type Repository interface {
	Save(shortUrl ShortUrl) (ShortUrl, error)
	FindByBackHalf(backHalf string) (ShortUrl, error)
	Update(shortUrl ShortUrl) (ShortUrl, error)
	Delete(ID int) (ShortUrl, error)
	FindAll() ([]ShortUrl, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(shortUrl ShortUrl) (ShortUrl, error) {
	err := r.db.Create(&shortUrl).Error

	if err != nil {
		return shortUrl, err
	}

	return shortUrl, nil
}

func (r *repository) FindByBackHalf(backHalf string) (ShortUrl, error) {
	var shortUrl ShortUrl
	err := r.db.Where("back_half = ?", backHalf).Find(&shortUrl).Error

	if err != nil {
		return shortUrl, err
	}

	return shortUrl, nil
}

func (r *repository) Update(shortUrl ShortUrl) (ShortUrl, error) {
	err := r.db.Save(&shortUrl).Error

	if err != nil {
		return shortUrl, err
	}

	return shortUrl, nil
}

func (r *repository) Delete(ID int) (ShortUrl, error) {
	var shortUrl ShortUrl
	err := r.db.Where("id = ?", ID).Delete(&shortUrl).Error

	if err != nil {
		return shortUrl, err
	}

	return shortUrl, nil
}

func (r *repository) FindAll() ([]ShortUrl, error) {
	var shortUrls []ShortUrl

	err := r.db.Find(&shortUrls).Error

	if err != nil {
		return shortUrls, err
	}

	return shortUrls, nil
}
