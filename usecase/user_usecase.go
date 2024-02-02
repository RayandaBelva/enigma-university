package usecase

import (
	"fmt"

	"github.com/eulbyvan/enigma-university/model"
	"github.com/eulbyvan/enigma-university/repository"
)

type UserUseCase interface {
	// create
	Registration(user model.User) error

	// read
	FindById(id string) (model.User, error)
	GetAllUsers() ([]model.User, error)

	// update
	UpdateById(id string, updatedUser model.User) error

	// delete
	RemoveById(id string) error
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	user, err := u.repo.GetById(id)

	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	return user, err
}

func (u *userUseCase) GetAllUsers() ([]model.User, error) {
	
	users, err := u.repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}

	return users, nil
}


func (u *userUseCase) Registration(user model.User) error {

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return fmt.Errorf("incomplete user data")
	}

	err := u.repo.Create(user)
	if err != nil {
		return fmt.Errorf("failed to register user: %v", err)
	}

	return nil
}

func (u *userUseCase) UpdateById(id string, updatedUser model.User) error {

	err := u.repo.UpdateById(id, updatedUser)
	if err != nil {
		return fmt.Errorf("failed to update user with ID %s: %v", id, err)
	}

	return nil
}

func (u *userUseCase) RemoveById(id string) error {

	err := u.repo.DeleteById(id)
	if err != nil {
		return fmt.Errorf("failed to delete user with ID %s: %v", id, err)
	}

	return nil
}

// constructor
func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
