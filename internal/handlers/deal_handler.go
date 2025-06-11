package handlers

import (
	"net/http"
	"encoding/json"
	"turcompany/internal/models"
	"turcompany/internal/services"
	"github.com/gorilla/mux"
)
type DealHandler struct {
	Service *services.DealService
}
func NewDealHandler(service *services.DealService) *DealHandler {
	return &DealHandler{Service: service}
}

func (h *DealHandler) Create(w http.ResponseWriter, r*http.Request){
	var deal models.Deal
	if err:= json.NewDecoder(r.Body).Decode(&deal); err!= nil {
		http.Error(w, err.Error(),http.StatusBadRequest)
		return
	}
	if err:=h.Service.Create(&deal); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *DealHandler) Update(w http.ResponseWriter, r *http.Request) {
	id:= mux.Vars(r)["id"]
	var deal models.Deal
	if err := json.NewDecoder(r.Body).Decode(&deal); err != nil{
		http.Error(w, err.Error(),http.StatusBadRequest)
		return
		}
	deal.ID =id
	if err:=h.Service.Update(&deal); err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *DealHandler) GetByID(w http.ResponseWriter, r*http.Request){
	id:=mux.Vars(r)["id"]
	deal,err:=h.Service.GetByID(id)
	if err != nil{
		http.Error(w,"Deal not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(deal)
}
func (h *DealHandler)Delete(w http.ResponseWriter,r *http.Request){
	id:= mux.Vars(r)["id"]
	if err:= h.Service.Delete(id); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}