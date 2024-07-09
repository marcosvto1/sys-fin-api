package handlers

import (
	"net/http"

	"github.com/go-chi/render"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg/errorable"
)

type DashboardController struct {
	GetChartTransactionByCategoryUseCase usecase.IChartTransactionByCategory
	GetChartTransactionByTypeUseCase     usecase.IChartTransactionByType
}

func NewDashboardController(getChartTransactionByCategoryUC usecase.IChartTransactionByCategory, getChartTransactionByTypeUC usecase.IChartTransactionByType) *DashboardController {
	return &DashboardController{
		GetChartTransactionByCategoryUseCase: getChartTransactionByCategoryUC,
		GetChartTransactionByTypeUseCase:     getChartTransactionByTypeUC,
	}
}

func (h *DashboardController) GetChartTransactionByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.GetChartTransactionByCategoryUseCase.Execute(dtos.GetChartTransactionByCategoryInput{
		Month: r.URL.Query().Get("month"),
		Year:  r.URL.Query().Get("year"),
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, res)
}

func (h *DashboardController) GetChartTransactionByTypeHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.GetChartTransactionByTypeUseCase.Execute(dtos.GetChartTransactionByTypeInput{
		Year: r.URL.Query().Get("year"),
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, res)
}
