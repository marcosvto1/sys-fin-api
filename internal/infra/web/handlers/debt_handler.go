package handlers

import (
	"encoding/json"
	"net/http"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"

		"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"

		"gitlab.com/marcosvto/sys-fin-api/pkg/errorable"
		"errors"
)

type DebtHandler struct {
	CreateDebtUseCase usecase.ICreateDebt
}

func NewDebtHandler(createDebtUseCase usecase.ICreateDebt) *DebtHandler {
	return &DebtHandler{
		CreateDebtUseCase: createDebtUseCase,
	}
}

func (h *DebtHandler) CreateDebt(w http.ResponseWriter, r *http.Request) {
	var createDebtDTO dtos.CreateDebtInput
	if err := json.NewDecoder(r.Body).Decode(&createDebtDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	if createDebtDTO.Amount <= 0 {
		http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)

		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{errors.New(errorable.CATEGORY_MUST_BE_GREATER_THAN_ZERO)}))

		return
	}

	if createDebtDTO.CategoryId <= 0 {
		// errr := http.Error(w, "Category ID must be greater than zero", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{errors.New(errorable.CATEGORY_MUST_BE_GREATER_THAN_ZERO)}))

		return
	}


	createdDebt, err := h.CreateDebtUseCase.Execute(createDebtDTO)

	if err != nil {
		log.Error("Error creating debt: ", err)
		http.Error(w, "Failed to create debt", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdDebt)
}	