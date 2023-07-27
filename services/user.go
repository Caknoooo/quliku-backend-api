package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDTO dto.UserCreateDTO) (entities.User, error)
	GetAllUser(ctx context.Context) ([]entities.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (entities.User, error)
	CheckUserEmail(ctx context.Context, email string) (bool, error)
	CheckUserUsername(ctx context.Context, username string) (bool, error)
	UpdateUser(ctx context.Context, userDTO dto.UserUpdateDTO) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	VerifyLogin(ctx context.Context, userDTO dto.UserLoginDTO) (bool, error)
	VerifyEmail(ctx context.Context, userVerificationDTO dto.UserVerificationDTO) (bool, error)
	ResendVerificationCode(ctx context.Context, userVerificationDTO dto.ResendVerificationCode) (bool, error)
	ResendFailedLoginNotVerified(ctx context.Context, email string) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
	userVeritificationRepository repository.UserVerificationRepository
}

func NewUserService(ur repository.UserRepository, uv repository.UserVerificationRepository) UserService {
	return &userService{
		userRepository: ur,
		userVeritificationRepository: uv,
	}
}

func (us *userService) RegisterUser(ctx context.Context, userDTO dto.UserCreateDTO) (entities.User, error) {
	user := entities.User{}

	if userDTO.Password != userDTO.ConfirmPassword {
		return entities.User{}, dto.ErrPasswordNotMatch
	}	

	err := smapping.FillStruct(&user, smapping.MapFields(userDTO))
	user.Role = "user"
	if err != nil {
		return entities.User{}, err
	}
	user, err = us.userRepository.RegisterUser(ctx, user)
	if err != nil {
		return entities.User{}, err
	}

	// Send verification email
	draftEmail, err := utils.MakeVerificationEmail(user.Email)
	if err != nil {
		return entities.User{}, err
	}

	// Expired verification code in 1 hour
	_ = us.userVeritificationRepository.Create(user.ID, draftEmail["code"], time.Now().Add(time.Minute * 3))

	err = utils.SendMail(user.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (us *userService) VerifyEmail(ctx context.Context, userVerificationDTO dto.UserVerificationDTO) (bool, error) {
	userVerification, err := us.userVeritificationRepository.Check(userVerificationDTO.UserID)
	if err != nil {
		return false, err
	}

	if userVerification.ReceiveCode != userVerificationDTO.SendCode {
		return false, dto.ErrorVerificationCodeNotMatch
	}

	if userVerification.ExpiredAt.Before(time.Now()) {
		return false, dto.ErrorExpiredVerificationCode
	}

	if userVerification.IsActive {
		return false, dto.ErrorUserVerificationCodeAlreadyUsed
	}

	if err := us.userVeritificationRepository.SendCode(userVerification.UserID, userVerificationDTO.SendCode); err != nil {
		return false, err
	}

	return true, nil
}

func (us *userService) ResendVerificationCode(ctx context.Context, userVerificationDTO dto.ResendVerificationCode) (bool, error) {
	userVerification, err := us.userVeritificationRepository.Check(userVerificationDTO.UserID)
	if err != nil {
		return false, err
	}

	if userVerification.ExpiredAt.After(time.Now()) {
		return false, dto.ErrorNotExpiredVerificationCode
	}

	if userVerification.IsActive {
		return false, dto.ErrorUserVerificationCodeAlreadyUsed
	}

	user, err := us.userRepository.GetUserByID(ctx, userVerification.UserID)
	if err != nil {
		return false, err
	}

	// Send verification email
	draftEmail, err := utils.MakeVerificationEmail(user.Email)
	if err != nil {
		return false, err
	}

	_ = us.userVeritificationRepository.Delete(userVerification.UserID)

	// Expired verification code in 1 hour
	_ = us.userVeritificationRepository.Create(userVerification.UserID, draftEmail["code"], time.Now().Add(time.Minute * 3))

	err = utils.SendMail(user.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return false, err
	}

	return true, nil
}

func (us *userService) GetAllUser(ctx context.Context) ([]entities.User, error) {
	return us.userRepository.GetAllUser(ctx)
}

func (us *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error) {
	return us.userRepository.GetUserByID(ctx, userID)
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	return us.userRepository.GetUserByEmail(ctx, email)
}

func (us *userService) GetUserByUsername(ctx context.Context, username string) (entities.User, error) {
	return us.userRepository.GetUserByUsername(ctx, username)
}

func (us *userService) CheckUserEmail(ctx context.Context, email string) (bool, error) {
	mail, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if mail.Email == "" {
		return false, err
	}
	return true, nil
}

func (us *userService) CheckUserUsername(ctx context.Context, username string) (bool, error) {
	user, err := us.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	if user.Username == "" {
		return false, err
	}
	return true, nil
}

func (us *userService) UpdateUser(ctx context.Context, userDTO dto.UserUpdateDTO) error {
	user := entities.User{}
	if err := smapping.FillStruct(&user, smapping.MapFields(userDTO)); err != nil {
		return nil
	}
	return us.userRepository.UpdateUser(ctx, user)
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return us.userRepository.DeleteUser(ctx, userID)
}

func (us *userService) VerifyLogin(ctx context.Context, userDTO dto.UserLoginDTO) (bool, error) {
	if userDTO.Email != "" {
		user, err := us.userRepository.GetUserByEmail(ctx, userDTO.Email)
		if err != nil {
			return false, err
		}

		if user.Email == "" {
			return false, dto.ErrorUserNotFound
		}

		if checkPassword, err := helpers.CheckPassword(user.Password, []byte(userDTO.Password)); !checkPassword {
			return false, err
		}

		return true, nil
	} else if userDTO.Username != "" {
		user, err := us.userRepository.GetUserByUsername(ctx, userDTO.Username)
		if err != nil {
			return false, err
		}
		fmt.Println(user.Username, user.Password)

		if user.Username == "" {
			return false, dto.ErrorUserNotFound
		}

		fmt.Println(user.Username, user.Password)
		if checkPassword, err := helpers.CheckPassword(user.Password, []byte(userDTO.Password)); !checkPassword {
			return false, err
		}

		return true, nil
	}

	return false, dto.ErrorUserNotFound
}

func (us *userService) ResendFailedLoginNotVerified(ctx context.Context, email string) (bool, error) {
	mail, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if mail.IsVerified {
		return false, dto.ErrorUserAlreadyActive
	}

	userVerification, err := us.userVeritificationRepository.Check(mail.ID)
	if err != nil {
		return false, err
	}

	if userVerification.ExpiredAt.After(time.Now()) {
		return false, dto.ErrorNotExpiredVerificationCode
	}

	if userVerification.IsActive {
		return false, dto.ErrorUserVerificationCodeAlreadyUsed
	}

	draftEmail, err := utils.MakeVerificationEmail(mail.Email)
	if err != nil {
		return false, err
	}

	_ = us.userVeritificationRepository.Delete(mail.ID)

	_ = us.userVeritificationRepository.Create(mail.ID, draftEmail["code"], time.Now().Add(time.Minute * 3))

	err = utils.SendMail(mail.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return false, err
	}

	return true, nil
}