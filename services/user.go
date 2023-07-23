package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
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
	Verify(ctx context.Context, userDTO dto.UserLoginDTO) (bool, error)

	// Mail
	// MakeVerificationEmail()
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
	draftEmail, err := MakeVerificationEmail(user.Email)
	if err != nil {
		return entities.User{}, err
	}

	// Expired verification code in 1 hour
	_ = us.userVeritificationRepository.Create(user.ID, draftEmail["code"], time.Now().Add(time.Hour))

	err = utils.SendMail(user.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func MakeVerificationEmail(receiverEmail string) (map[string]string, error) {
	token := EncodeToString(6)
	if token == "" {
		return nil, dto.ErrorFailedGenerateVerificationCode
	}

	draftEmail := map[string]string{}
	draftEmail["subject"] = "Quliku - Email Verification"
	draftEmail["code"] = token
	draftEmail["body"] = fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Email Verification - Quliku</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					line-height: 1.6;
				}
				.code-container {
					text-align: center;
				}
				.code {
					font-size: 24px;
					background-color: #f1f1f1;
					padding: 10px;
					display: inline-block;
				}
				.note {
					font-size: 14px;
					color: #888;
				}
			</style>
		</head>
		<body>
			<p>Hi, %s! Thanks for registering an account on Quliku.App.</p>
			<p>Before you can access your account, you need to verify your email address by clicking the link below:</p>
			<div class="code-container">
				<p class="code">%s</p>
			</div>
			<p class="note">Please note that this link will expire after 1 hour.</p>
			<p>If you did not sign up for an account on Quliku, you can safely ignore this email.</p>
		</body>
		</html>
	`, receiverEmail, token)

	return draftEmail, nil
}

func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, max)
	n, _ := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return dto.ErrorFailedGenerateVerificationCode.Error()
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i]) % len(table)]
	}
 
	return string(b)
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

func (us *userService) Verify(ctx context.Context, userDTO dto.UserLoginDTO) (bool, error) {
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

