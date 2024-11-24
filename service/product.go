package service

import (
	"math"

	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/repository"
)

type ServiceProduct struct {
	Repo *repository.RepoProduct
}

func NewServiceProduct(repo *repository.RepoProduct) *ServiceProduct {
	return &ServiceProduct{repo}
}

func (s *ServiceProduct) GetAll(qp model.QueryProduct) (model.Response, error) {
	products, totalItems, err := s.Repo.FindAll(qp)
	if err != nil {
		return model.Response{}, err
	}
	if qp.Page == 0 {
		qp.Page = 1
	}
	response := model.Response{
		Page:       qp.Page,
		Limit:      6,
		TotalItem:  totalItems,
		TotalPages: int(math.Ceil(float64(totalItems) / float64(6))),
		Data:       products,
	}
	return response, nil
}
func (s *ServiceProduct) GetById(product *model.Product) error {
	err := s.Repo.FindById(product)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceProduct) GetBanner() ([]model.Banner, error) {
	banners, err := s.Repo.FindBanner()
	if err != nil {
		return []model.Banner{}, err
	}
	return banners, nil
}
func (s *ServiceProduct) GetPromo() ([]model.ShowcaseProduct, error) {
	promos, err := s.Repo.FindPromo()
	if err != nil {
		return []model.ShowcaseProduct{}, err
	}
	return promos, nil
}
func (s *ServiceProduct) GetRecommend() ([]model.ShowcaseProduct, error) {
	recommends, err := s.Repo.FindRecommend()
	if err != nil {
		return []model.ShowcaseProduct{}, err
	}
	return recommends, nil
}
func (s *ServiceProduct) GetAllCategory() ([]model.Category, error) {
	categories, err := s.Repo.FindAllCategory()
	if err != nil {
		return []model.Category{}, err
	}
	return categories, nil
}
