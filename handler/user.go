package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/service"
	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
)

type HandlerUser struct {
	Service *service.ServiceUser
}

func NewHandlerUser(service *service.ServiceUser) *HandlerUser {
	return &HandlerUser{service}
}

func (h *HandlerUser) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, "Invalid Request Body")
		json.NewEncoder(w).Encode(response)
		return
	}
	errValidate, err := utils.ValidateInput(user)
	if err != nil {
		response := utils.SetResponse(w, model.Response{Data: errValidate}, http.StatusBadRequest, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	err = h.Service.Register(&user)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusCreated, "Register Account Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerUser) Login(w http.ResponseWriter, r *http.Request) {
	var form model.Login
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, "Invalid Request Body")
		json.NewEncoder(w).Encode(response)
		return
	}
	errValidate, err := utils.ValidateInput(form)
	if err != nil {
		response := utils.SetResponse(w, model.Response{Data: errValidate}, http.StatusBadRequest, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	err = h.Service.Login(&user, form)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	utils.SetCookie(w, "id", strconv.Itoa(user.Id))
	utils.SetCookie(w, "token", *user.Token)
	response := utils.SetResponse(w, model.Response{}, http.StatusOK, "Login Success")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerUser) Logout(w http.ResponseWriter, r *http.Request) {
	utils.SetCookie(w, "id", "")
	utils.SetCookie(w, "token", "")
	response := utils.SetResponse(w, model.Response{}, http.StatusOK, "Logout Success")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Cookie("id")
	token, _ := r.Cookie("token")
	user := model.User{
		Id:    utils.ToInt(id.Value),
		Token: &token.Value,
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, "Invalid Request Body")
		json.NewEncoder(w).Encode(response)
		return
	}
	errValidate, err := utils.ValidateInput(user)
	if err != nil {
		response := utils.SetResponse(w, model.Response{Data: errValidate}, http.StatusBadRequest, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	err = h.Service.UpdateUser(&user)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{}, http.StatusOK, "Update Account Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerUser) GetAddresses(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Cookie("id")
	data, err := h.Service.GetAddresses(utils.ToInt(userId.Value))
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: data}, http.StatusOK, "Get Addresses Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerUser) AddAddress(w http.ResponseWriter, r *http.Request) {
	var address model.Address
	userId, _ := r.Cookie("id")
	address.UserId = utils.ToInt(userId.Value)
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, "Invalid Request Body")
		json.NewEncoder(w).Encode(response)
		return
	}
	err := h.Service.AddAddress(&address)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: address}, http.StatusCreated, "Add Address Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerUser) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var address model.Address
	userId, _ := r.Cookie("id")
	address.Id = utils.ToInt(chi.URLParam(r, "id"))
	address.UserId = utils.ToInt(userId.Value)
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, "Invalid Request Body")
		json.NewEncoder(w).Encode(response)
		return
	}
	err := h.Service.UpdateAddress(&address)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: address}, http.StatusOK, "Update Address Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerUser) SetDefaultAddress(w http.ResponseWriter, r *http.Request) {
	var address model.Address
	userId, _ := r.Cookie("id")
	address.Id = utils.ToInt(chi.URLParam(r, "id"))
	address.UserId = utils.ToInt(userId.Value)
	err := h.Service.SetDefaultAddress(&address)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: address}, http.StatusOK, "Set Default Address Succes")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerUser) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	var address model.Address
	userId, _ := r.Cookie("id")
	address.Id = utils.ToInt(chi.URLParam(r, "id"))
	address.UserId = utils.ToInt(userId.Value)
	err := h.Service.DeleteAddress(&address)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: address}, http.StatusOK, "Delete Address Succes")
	json.NewEncoder(w).Encode(response)
}
