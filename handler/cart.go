package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/service"
	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
)

type HandlerCart struct {
	Service *service.ServiceCart
}

func NewHandlerCart(service *service.ServiceCart) *HandlerCart {
	return &HandlerCart{service}
}

func (h *HandlerCart) GetAll(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Cookie("id")
	carts, err := h.Service.GetAll(utils.ToInt(userId.Value))
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: carts}, http.StatusOK, "Get All Cart Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerCart) AddCart(w http.ResponseWriter, r *http.Request) {
	var cart model.Cart
	userId, _ := r.Cookie("id")
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, "Invalid Body Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	cart.UserId = utils.ToInt(userId.Value)
	err := h.Service.AddCart(&cart)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusCreated, "Add Cart Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerCart) UpdateIncrementQty(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Cookie("id")
	id := utils.ToInt(chi.URLParam(r, "productVariantId"))
	err := h.Service.UpdateIncrementQty(id, utils.ToInt(userId.Value))
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusCreated, "Increase Qty Cart Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerCart) UpdateDecrementQty(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Cookie("id")
	id := utils.ToInt(chi.URLParam(r, "productVariantId"))
	err := h.Service.UpdateDecrementQty(id, utils.ToInt(userId.Value))
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusCreated, "Decrease Qty Cart Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerCart) DeleteCart(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Cookie("id")
	id := utils.ToInt(chi.URLParam(r, "productVariantId"))
	err := h.Service.DeleteCart(id, utils.ToInt(userId.Value))
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusOK, "Delete Cart Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerCart) GetTotal(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Cookie("id")
	totalCart, err := h.Service.GetTotal(utils.ToInt(userId.Value))
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: totalCart}, http.StatusOK, "Delete Cart Succes")
	json.NewEncoder(w).Encode(response)
}
