package repository

import (
	"time"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserVerificationRepository interface {
	Create(UserID uuid.UUID, receiveCode string, expiredAt time.Time) error
	SendCode(UserID uuid.UUID, sendCode string) error
	Check(UserID uuid.UUID) (entities.UserVerification, error)
}

type userVerificationRepository struct {
	db *gorm.DB
}

func NewUserVerificationRepository(db *gorm.DB) *userVerificationRepository {
	return &userVerificationRepository{
		db: db,
	}
}

func (u *userVerificationRepository) Create(UserID uuid.UUID, receiveCode string, expiredAt time.Time) (error) {
	userVerification := entities.UserVerification{
		UserID: UserID,
		ReceiveCode: receiveCode,
		ExpiredAt: expiredAt,
	}

	if err := u.db.Create(&userVerification).Error; err != nil {
		return err
	}

	return nil
}

func(u *userVerificationRepository) SendCode(UserID uuid.UUID, sendCode string) error {
	if err := u.db.Model(&entities.UserVerification{}).Where("user_id = ?", UserID).Update("send_code", sendCode).Error; err != nil {
		return err
	}

	if err := u.db.Model(&entities.UserVerification{}).Where("user_id = ?", UserID).Update("is_active", true).Error; err != nil {
		return err
	}

	if err := u.db.Model(&entities.User{}).Where("id = ?", UserID).Update("is_verified", true).Error; err != nil {
		return err
	}

	return nil
}

func(u *userVerificationRepository) Check(UserID uuid.UUID) (entities.UserVerification, error) {
	var userVerification entities.UserVerification

	if err := u.db.Model(&entities.UserVerification{}).Where("user_id = ?", UserID).First(&userVerification).Error; err != nil {
		return entities.UserVerification{}, err
	}

	return userVerification, nil
}