package shorter

import (
	"errors"
	"id/projects/url-shortener/helper"
)

type Service interface {
	GenerateShorterUrl(input UrlInput) (ShortUrl, error)
	SaveShorterUrl(input ShorterUrlInput) (ShortUrl, error)
	FindShorterUrl(input string) (ShortUrl, error)
	FindAllShorterUrl() ([]ShortUrl, error)
	UpdateShorterUrl(input ShorterUrlInput) (ShortUrl, error)
	DeleteShorterUrl(ID int) (ShortUrl, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GenerateShorterUrl(input UrlInput) (ShortUrl, error) {
	shortUrl := ShortUrl{}
	shortUrl.Title = input.Title
	shortUrl.DestinationUrl = input.DestinationUrl
	shortUrl.Tags = input.Tags

	backHalf := helper.GenerateRandom()
	if backHalf == "" {
		return shortUrl, errors.New("Error when create unique back half")
	}
	shortUrl.BackHalf = backHalf

	return shortUrl, nil
}

func (s *service) SaveShorterUrl(input ShorterUrlInput) (ShortUrl, error) {
	existUrl, err := s.repository.FindByBackHalf(input.BackHalf)

	if err != nil {
		return existUrl, err
	}

	if existUrl.ID != 0 {
		return existUrl, errors.New("Back half already exist")
	}

	shortUrl := ShortUrl{}
	shortUrl.Title = input.Title
	shortUrl.DestinationUrl = input.DestinationUrl
	shortUrl.Tags = input.Tags
	shortUrl.BackHalf = input.BackHalf

	newShortUrl, err := s.repository.Save(shortUrl)

	if err != nil {
		return newShortUrl, err
	}

	return newShortUrl, nil
}

func (s *service) FindShorterUrl(input string) (ShortUrl, error) {
	shortUrl, err := s.repository.FindByBackHalf(input)

	if err != nil {
		return shortUrl, err
	}

	if shortUrl.ID == 0 {
		return shortUrl, errors.New("No back half found")
	}

	return shortUrl, nil
}

func (s *service) FindAllShorterUrl() ([]ShortUrl, error) {
	shortUrls, err := s.repository.FindAll()

	if err != nil {
		return shortUrls, err
	}

	return shortUrls, nil
}

func (s *service) UpdateShorterUrl(input ShorterUrlInput) (ShortUrl, error) {
	existUrl, err := s.repository.FindByBackHalf(input.BackHalf)

	if err != nil {
		return existUrl, err
	}

	shortUrl := ShortUrl{}
	shortUrl.Title = input.Title
	shortUrl.DestinationUrl = input.DestinationUrl
	shortUrl.Tags = input.Tags
	shortUrl.BackHalf = input.BackHalf

	updateShorterUrl, err := s.repository.Update(shortUrl)

	if err != nil {
		return updateShorterUrl, err
	}

	return updateShorterUrl, nil
}

func (s *service) DeleteShorterUrl(ID int) (ShortUrl, error) {
	shortUrl, err := s.repository.Delete(ID)

	if err != nil {
		return shortUrl, err
	}

	return shortUrl, nil
}
