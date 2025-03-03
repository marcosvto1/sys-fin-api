package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
)

type SubscriptionHandler struct {
	FindSubscriptionUC   *usecase.FindSubscriptionUsecase
	CreateSubscriptionUC *usecase.CreateSubscriptionUsecase
	DeleteSubscriptionUC *usecase.DeleteSubscriptionUseCase
	UpdateSubscriptionUC *usecase.UpdateSubscriptionUsecase
}

func NewSubscriptionHandler(
	findSubscription *usecase.FindSubscriptionUsecase,
	createSubscription *usecase.CreateSubscriptionUsecase,
	deleteSubscription *usecase.DeleteSubscriptionUseCase,
	updateSubscription *usecase.UpdateSubscriptionUsecase,
) *SubscriptionHandler {
	return &SubscriptionHandler{
		FindSubscriptionUC:   findSubscription,
		CreateSubscriptionUC: createSubscription,
		DeleteSubscriptionUC: deleteSubscription,
		UpdateSubscriptionUC: updateSubscription,
	}
}

func (h *SubscriptionHandler) FindAllHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *SubscriptionHandler) FindByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subscription, err := h.FindSubscriptionUC.FindById(intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subscription)
}

func (h *SubscriptionHandler) DeleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.FindSubscriptionUC.FindById(intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.DeleteSubscriptionUC.Execute(intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *SubscriptionHandler) UpdateByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.FindSubscriptionUC.FindById(intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var input dtos.UpdateSubscriptionInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.UpdateSubscriptionUC.Execute(intId, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}
