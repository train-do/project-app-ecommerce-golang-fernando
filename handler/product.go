package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/service"
	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
)

type HandlerProduct struct {
	Service *service.ServiceProduct
}

func NewHandlerProduct(service *service.ServiceProduct) *HandlerProduct {
	return &HandlerProduct{service}
}

func (h *HandlerProduct) GetAll(w http.ResponseWriter, r *http.Request) {
	var qp model.QueryProduct
	qp.Page = utils.ToInt(r.URL.Query().Get("page"))
	qp.Name = r.URL.Query().Get("name")
	qp.Category = r.URL.Query().Get("category")
	qp.IsBestSelling = utils.ToBool(r.URL.Query().Get("best_selling"))
	data, err := h.Service.GetAll(qp)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, data, http.StatusOK, "Get All Product Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerProduct) GetById(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	product.Id = utils.ToInt(chi.URLParam(r, "id"))
	err := h.Service.GetById(&product)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: product}, http.StatusOK, "Get Product Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerProduct) GetBanner(w http.ResponseWriter, r *http.Request) {
	banners, err := h.Service.GetBanner()
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: banners}, http.StatusOK, "Get All Product Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerProduct) GetPromo(w http.ResponseWriter, r *http.Request) {
	promos, err := h.Service.GetPromo()
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: promos}, http.StatusOK, "Get All Promo Product Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerProduct) GetRecommend(w http.ResponseWriter, r *http.Request) {
	recommends, err := h.Service.GetRecommend()
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: recommends}, http.StatusOK, "Get All Recommend Product Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerProduct) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := h.Service.GetAllCategory()
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: categories}, http.StatusOK, "Get All Category Succes")
	json.NewEncoder(w).Encode(response)
}
