package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-adv-api/pkg/errorable"
)

type TransactionController struct {
	CreateTransactionUseCase usecase.ICreateTransaction
	FindTransactionUseCase   usecase.IFindTransaction
	FindOneTrasactionUseCase usecase.IFindOneTransaction
	DeleteTransactionUseCase usecase.IDeleteTransaction
}

func NewTransactionController(
	createTransactionUC usecase.ICreateTransaction,
	findTransactionUC usecase.IFindTransaction,
	findOneTransactionUC usecase.IFindOneTransaction,
	deleteTransactionUC usecase.IDeleteTransaction,
) *TransactionController {
	return &TransactionController{
		CreateTransactionUseCase: createTransactionUC,
		FindTransactionUseCase:   findTransactionUC,
		FindOneTrasactionUseCase: findOneTransactionUC,
		DeleteTransactionUseCase: deleteTransactionUC,
	}
}

func (controller *TransactionController) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
		return
	}

	defer r.Body.Close()

	input := dtos.CreateTransactionInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
		return
	}

	output, err := controller.CreateTransactionUseCase.Execute(input)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, output)
}

func (controller *TransactionController) FindTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var errors []errorable.CtxError

	pId := r.URL.Query().Get("id")
	id := -1
	if pId != "" {
		idConv, err := strconv.Atoi(pId)
		if err != nil {
			errors = append(errors, *errorable.New(errorable.INVALID_VALUE_FIELD))
		}
		id = idConv
	}

	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	var categoryId int = -1
	cId := r.URL.Query().Get("category_id")
	if cId != "" {
		categoryIdInteger, err := strconv.Atoi(cId)
		if err != nil {
			log.Println(err)
			errors = append(errors, *errorable.New(errorable.INVALID_VALUE_FIELD))
		}
		categoryId = categoryIdInteger
	}

	var walletId int = -1
	wId := r.URL.Query().Get("wallet_id")
	if wId != "" {
		walletIdInteger, err := strconv.Atoi(wId)
		if err != nil {
			errors = append(errors, *errorable.New(errorable.INVALID_VALUE_FIELD))
		}
		walletId = walletIdInteger
	}

	pageNumber, err := strconv.Atoi(r.URL.Query().Get("page_number"))
	if err != nil {
		log.Println(err)
		errors = append(errors, *errorable.New(errorable.INVALID_VALUE_FIELD))
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		errors = append(errors, *errorable.New(errorable.INVALID_VALUE_FIELD))
	}

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.NewHttpError(http.StatusBadRequest, errors))
		return
	}

	input := dtos.FindTransactionInput{
		Id:         id,
		Month:      month,
		Year:       year,
		PageNumber: pageNumber,
		PageSize:   pageSize,
		CategoryId: categoryId,
		WalletId:   walletId,
	}

	output, err := controller.FindTransactionUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, errorable.NewHttpError(http.StatusInternalServerError, []errorable.CtxError{
			*errorable.New(errorable.INTERNAL_ERROR),
		}))
		return
	}

	json.NewEncoder(w).Encode(output)
}

func (controller *TransactionController) FindOneTransactionHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	if idString == "" {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{errors.New("path params :id does not founded")}))
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{errors.New("invalid param :id")}))
		return
	}

	output, err := controller.FindOneTrasactionUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, errorable.HttpResponse(http.StatusNotFound, []error{errors.New(errorable.NOT_FOUND_REGISTER)}))
		return
	}

	render.JSON(w, r, output)
}

func (controller *TransactionController) DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	if idString == "" {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{errors.New("path params :id does not founded")}))
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{errors.New("invalid param :id")}))
		return
	}

	err = controller.DeleteTransactionUseCase.Execute(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadGateway)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadGateway, []error{err}))
		return
	}
}
