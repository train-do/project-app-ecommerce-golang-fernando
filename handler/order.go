package handler

import (
	"encoding/json"
	"net/http"

	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/service"
	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
)

type HandlerOrder struct {
	Service *service.ServiceOrder
}

func NewHandlerOrder(service *service.ServiceOrder) *HandlerOrder {
	return &HandlerOrder{service}
}

func (h *HandlerOrder) Create(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Cookie("id")
	var orderProduct model.OrderProduct
	if err := json.NewDecoder(r.Body).Decode(&orderProduct); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, "Invalid Request Body")
		json.NewEncoder(w).Encode(response)
		return
	}
	err := h.Service.Create(utils.ToInt(userId.Value), orderProduct)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusCreated, "Create Order Succes")
	json.NewEncoder(w).Encode(response)
}
