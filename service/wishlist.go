package service

import (
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/repository"
)

type ServiceWishlist struct {
	Repo *repository.RepoWishlist
}

func NewServiceWishlist(repo *repository.RepoWishlist) *ServiceWishlist {
	return &ServiceWishlist{repo}
}

func (s *ServiceWishlist) AddWishlist(wishlist *model.Wishlist) error {
	err := s.Repo.Insert(wishlist)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceWishlist) DeleteWishlist(wishlist *model.Wishlist) error {
	err := s.Repo.Delete(wishlist)
	if err != nil {
		return err
	}
	return nil
}
