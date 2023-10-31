package services

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"github.com/php-redneck/volga-it-simbir-go/internal/entities"
	"github.com/php-redneck/volga-it-simbir-go/internal/repository"
	"github.com/php-redneck/volga-it-simbir-go/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository repository.UserRepository
}

func (s UserService) Total() int64 {
	var (
		count int64
		user  entities.User
	)

	database.Instance().Model(&user).Count(&count)

	return count
}

func (s UserService) FindByUsername(username string) (entities.User, error) {
	return s.Repository.FindBy("username", username)
}

func (s UserService) Index(start, count int) []entities.User {
	return repository.UserRepository{}.Index(start, count)
}

func (s UserService) Show(id int) (entities.User, error) {
	return repository.UserRepository{}.FindBy("id", id)
}

func (s UserService) Store(dto dto.CreateUserDTO) (entities.User, error) {
	user, err := s.FindByUsername(dto.Username)

	if err == nil {
		return user, errors.New(fmt.Sprintf("пользователь с ником %s существует", dto.Username))
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.MinCost)

	if err != nil {
		log.Error(err.Error())
	}

	user = entities.User{
		Username: dto.Username,
		Password: string(bytes),
		IsAdmin:  dto.IsAdmin,
		Balance:  dto.Balance,
	}

	result := database.Instance().Create(&user)

	if result.Error != nil {
		panic(result.Error)
	}

	return user, nil
}

func (s UserService) Edit(id int, dto dto.EditUserDTO) (entities.User, error) {
	user, err := s.Show(id)

	if err != nil {
		return user, errors.New("not_found")
	}

	user.Balance = dto.Balance

	user.IsAdmin = dto.IsAdmin

	if dto.Username != "" {
		_, err = s.FindByUsername(dto.Username)

		if err == nil {
			return user, errors.New(fmt.Sprintf("пользователь с ником %s существует", dto.Username))
		}

		user.Username = dto.Username
	}

	if dto.Password != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.MinCost)

		if err != nil {
			panic(err.Error())
		}

		user.Password = string(bytes)
	}

	res := database.Instance().Save(&user)

	return user, res.Error
}

func (s UserService) Destroy(id int) error {
	user, err := s.Show(id)

	if err != nil {
		return err
	}

	database.Instance().Delete(&user)

	return nil
}
