package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/service"
	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
)

type HandlerWishlist struct {
	Service *service.ServiceWishlist
}

func NewHandlerWishlist(service *service.ServiceWishlist) *HandlerWishlist {
	return &HandlerWishlist{service}
}

func (h *HandlerWishlist) AddWishlist(w http.ResponseWriter, r *http.Request) {
	var wishlist model.Wishlist
	userId, _ := r.Cookie("id")
	if err := json.NewDecoder(r.Body).Decode(&wishlist); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, "Invalid Body Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	wishlist.UserId = utils.ToInt(userId.Value)
	err := h.Service.AddWishlist(&wishlist)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusCreated, "Add Wishlist Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerWishlist) DeleteWishlist(w http.ResponseWriter, r *http.Request) {
	var wishlist model.Wishlist
	userId, _ := r.Cookie("id")
	wishlist.UserId = utils.ToInt(userId.Value)
	wishlist.Id = utils.ToInt(chi.URLParam(r, "id"))
	err := h.Service.DeleteWishlist(&wishlist)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusOK, "Delete Wishlist Succes")
	json.NewEncoder(w).Encode(response)
}
