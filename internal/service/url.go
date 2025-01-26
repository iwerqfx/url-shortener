package service

import (
	"github.com/iwerqfx/url-shortener/internal/model"
	"github.com/iwerqfx/url-shortener/internal/repository/sqlite"
	"github.com/iwerqfx/url-shortener/internal/util"
)

type URLService interface {
	Create(url string) (alias string, err error)
	GetByAlias(alias string) (model.URL, error)
}

type urlService struct {
	*Service
	repository sqlite.URLRepository
}

func NewURLService(service *Service, repository sqlite.URLRepository) URLService {
	return &urlService{
		Service:    service,
		repository: repository,
	}
}

func (s *urlService) Create(url string) (string, error) {
	alias, err := util.GenerateAlias()
	if err != nil {
		return "", model.ErrInternalServer
	}

	if err = s.repository.Create(url, alias); err != nil {
		return "", err
	}

	return alias, nil
}

func (s *urlService) GetByAlias(alias string) (model.URL, error) {
	url, err := s.repository.GetByAlias(alias)
	if err != nil {
		return model.URL{}, err
	}

	if err = s.repository.IncreaseViews(alias); err != nil {
		return model.URL{}, model.ErrInternalServer
	}

	return url, nil
}
