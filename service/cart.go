package service

import (
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/repository"
)

type ServiceCart struct {
	Repo *repository.RepoCart
}

func NewServiceCart(repo *repository.RepoCart) *ServiceCart {
	return &ServiceCart{repo}
}

func (s *ServiceCart) GetAll(userId int) ([]model.Cart, error) {
	carts, err := s.Repo.FindAll(userId)
	if err != nil {
		return []model.Cart{}, err
	}
	return carts, nil
}
func (s *ServiceCart) AddCart(Cart *model.Cart) error {
	err := s.Repo.Insert(Cart)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceCart) UpdateIncrementQty(id int, userId int) error {
	err := s.Repo.UpdateIncrementQty(id, userId)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceCart) UpdateDecrementQty(id int, userId int) error {
	err := s.Repo.UpdateDecrementQty(id, userId)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceCart) DeleteCart(id int, userId int) error {
	err := s.Repo.Delete(id, userId)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceCart) GetTotal(userId int) (int, error) {
	return s.Repo.FindTotal(userId)
}
