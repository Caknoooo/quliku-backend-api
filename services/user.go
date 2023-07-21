package services

import (
	"context"
	"fmt"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/repository"
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
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
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
	return us.userRepository.RegisterUser(ctx, user)
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

