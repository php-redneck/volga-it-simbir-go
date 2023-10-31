package services

import (
	"errors"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	UserService UserService
	JWTService  JWTService
}

func (s AccountService) SignIn(dto dto.SignInDTO) (string, error) {
	user, err := s.UserService.FindByUsername(dto.Login)

	if err != nil {
		return "", errors.New("пользователь отсутствует в системе")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if err != nil {
		return "", errors.New("неверный логин или пароль")
	}

	return s.JWTService.Make(int(user.ID), user.Username, user.IsAdmin), nil
}

func (s AccountService) SignUp(signInDto dto.SignUpDTO) (string, error) {
	_, err := s.UserService.FindByUsername(signInDto.Login)

	if err == nil {
		return "", errors.New("пользователь с таким ником присутствует в системе")
	}

	_, err = s.UserService.Store(dto.CreateUserDTO{
		Username: signInDto.Login,
		Password: signInDto.Password,
		IsAdmin:  false,
		Balance:  0,
	})

	if err != nil {
		panic(err)
	}

	return s.SignIn(dto.SignInDTO{Login: signInDto.Login, Password: signInDto.Password})
}

func (s AccountService) Update(userId int, updateAccountDTO dto.AccountUpdateDTO) (entities.User, error) {
	user, err := s.UserService.Show(userId)

	if err != nil {
		panic(err)
	}

	return s.UserService.Edit(userId, dto.EditUserDTO{
		Username: updateAccountDTO.Username,
		Password: updateAccountDTO.Password,
		IsAdmin:  user.IsAdmin,
		Balance:  user.Balance,
	})
}
