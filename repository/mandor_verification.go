package repository

import (
	"time"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MandorVerificationRepository interface {
	Create(mandorID uuid.UUID, receiveCode string, expiredAt time.Time) error
	Update(receiveCode string, expiredAt time.Time) error
	Delete(mandorID uuid.UUID) error
	SendCode(mandorID uuid.UUID, sendCode string) error
	Check(mandorID uuid.UUID) (entities.MandorVerification, error)
}

type mandorVerificationRepository struct {
	db *gorm.DB
}

func NewMandorVerificationRepository(db *gorm.DB) *mandorVerificationRepository {
	return &mandorVerificationRepository{
		db: db,
	}
}

func (m *mandorVerificationRepository) Create(mandorID uuid.UUID, receiveCode string, expiredAt time.Time) error {
	mandorVerification := entities.MandorVerification{
		MandorID:    mandorID,
		ReceiveCode: receiveCode,
		ExpiredAt:   expiredAt,
	}

	if err := m.db.Create(&mandorVerification).Error; err != nil {
		return err
	}

	return nil
}

func (m *mandorVerificationRepository) Update(receiveCode string, expiredAt time.Time) error {
	mandorVerification := entities.MandorVerification{
		ReceiveCode: receiveCode,
		ExpiredAt:   expiredAt,
	}

	if err := m.db.Model(&entities.MandorVerification{}).Updates(&mandorVerification).Error; err != nil {
		return err
	}

	return nil
}

func (m *mandorVerificationRepository) Delete(mandorID uuid.UUID) error {
	if err := m.db.Where("mandor_id", mandorID).Delete(&entities.MandorVerification{}).Error; err != nil {
		return err
	}

	return nil
}

func (m *mandorVerificationRepository) SendCode(mandorID uuid.UUID, sendCode string) error {
	mandorVerification := entities.MandorVerification{
		SendCode: sendCode,
		IsActive: true,
	}

	if err := m.db.Model(&entities.MandorVerification{}).Where("mandor_id = ?", mandorID).Updates(&mandorVerification).Error; err != nil {
		return err
	}

	if err := m.db.Model(&entities.Mandor{}).Where("id = ?", mandorID).Update("is_verified", true).Error; err != nil {
		return err
	}

	return nil
}

func (m *mandorVerificationRepository) Check(mandorID uuid.UUID) (entities.MandorVerification, error) {
	var mandorVerification entities.MandorVerification

	if err := m.db.Model(&entities.MandorVerification{}).Where("mandor_id = ?", mandorID).First(&mandorVerification).Error; err != nil {
		return entities.MandorVerification{}, err
	}

	return mandorVerification, nil
}
