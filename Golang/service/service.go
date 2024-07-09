package service

import (
	"mading/helper/logger"
	"mading/repository"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	repository *repository.Repository
	validate   *validator.Validate
	log        *logger.Log
}

func New(repo *repository.Repository, validate *validator.Validate) *Service {
	return &Service{
		repository: repo,
		validate:   validate,
		log:        logger.New("Service"),
	}
}
