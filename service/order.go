package service

import (
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/repository"
)

type ServiceOrder struct {
	Repo *repository.RepoOrder
}

func NewServiceOrder(repo *repository.RepoOrder) *ServiceOrder {
	return &ServiceOrder{repo}
}

func (s *ServiceOrder) Create(userId int, orderProduct model.OrderProduct) error {
	err := s.Repo.Insert(userId, orderProduct)
	if err != nil {
		return err
	}
	return nil
}
