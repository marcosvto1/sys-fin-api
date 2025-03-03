package handlers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
)

type SubscriptionHandler struct {
	FindSubscriptionUC   *usecase.FindSubscriptionUC
	CreateSubscriptionUC *usecase.CreateSubscriptionUC
}

func NewSubscriptionHandler(
	findSubscription *usecase.FindSubscriptionUC, createSubscription *usecase.CreateSubscriptionUC) *SubscriptionHandler {
	return &SubscriptionHandler{
		FindSubscriptionUC:   findSubscription,
		CreateSubscriptionUC: createSubscription,
	}
}

func (h *SubscriptionHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	subscriptions, err := h.FindSubscriptionUC.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subscriptions)
}

func (h *SubscriptionHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var input dtos.CreateSubscriptionInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subscription, err := h.CreateSubscriptionUC.Create(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subscription)
}
