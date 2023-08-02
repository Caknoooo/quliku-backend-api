package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type (
	MandorRatingService interface {
		CreateRatingMandor(ctx context.Context, mandorRating entities.MandorRating) (entities.MandorRating, error)
	}

	mandorRatingService struct {
		mandorRatingRepository repository.MandorRatingRepository
	}
)

func NewMandorRatingService(mandorRatingRepository repository.MandorRatingRepository) MandorRatingService {
	return &mandorRatingService{
		mandorRatingRepository: mandorRatingRepository,
	}
}

func (mrs *mandorRatingService) CreateRatingMandor(ctx context.Context, mandorRating entities.MandorRating) (entities.MandorRating, error) {
	var mandorRatingCreated entities.MandorRating

	mandorRatingCreated, err := mrs.mandorRatingRepository.CreateRatingMandor(ctx, mandorRating)
	if err != nil {
		return entities.MandorRating{}, err
	}

	return mandorRatingCreated, nil
}