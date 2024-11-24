package service

import (
	"github.com/google/uuid"
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/repository"
)

type ServiceUser struct {
	Repo *repository.RepoUser
}

func NewServiceUser(repo *repository.RepoUser) *ServiceUser {
	return &ServiceUser{repo}
}

func (s *ServiceUser) Register(user *model.User) error {
	err := s.Repo.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceUser) Login(user *model.User, form model.Login) error {
	user.Email = form.EmailOrPhone
	user.Phone = form.EmailOrPhone
	user.Password = form.Password
	user.Token = new(string) // karena inisiasi pertama nill, sehingga perlu inisiasi yang bukan nill
	*user.Token = uuid.New().String()
	err := s.Repo.Login(user)
	if err != nil {
		return err
	}
	user = &model.User{}
	return nil
}
func (s *ServiceUser) UpdateUser(user *model.User) error {
	err := s.Repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceUser) GetAddresses(userId int) ([]model.Address, error) {
	data, err := s.Repo.FindAddressesByUserId(userId)
	if err != nil {
		return []model.Address{}, err
	}
	return data, nil
}
func (s *ServiceUser) AddAddress(address *model.Address) error {
	err := s.Repo.InsertAddress(address)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceUser) UpdateAddress(address *model.Address) error {
	err := s.Repo.UpdateAddress(address)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceUser) SetDefaultAddress(address *model.Address) error {
	err := s.Repo.UpdateDefaultAddress(address)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceUser) DeleteAddress(address *model.Address) error {
	err := s.Repo.DeleteAddress(address)
	if err != nil {
		return err
	}
	return nil
}
