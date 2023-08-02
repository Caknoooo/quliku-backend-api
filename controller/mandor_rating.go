package controller

import (
	"github.com/Caknoooo/golang-clean_template/services"
)

type (
	MandorRatingController interface {
	}

	mandorRatingController struct {
		mandorRatingService services.MandorRatingService
	}
)

func NewMandorRatingController(mandorRatingService services.MandorRatingService) MandorRatingController {
	return &mandorRatingController{
		mandorRatingService: mandorRatingService,
	}
}