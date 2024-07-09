package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg/errorable"
)

type CategoryController struct {
	CreateCategoryUseCase usecase.ICreateCategory
	FindCategoriesUseCase usecase.IFindCategories
}

func NewCategoryController(createCategoryUC usecase.ICreateCategory, findCategoryUC usecase.IFindCategories) *CategoryController {
	return &CategoryController{
		CreateCategoryUseCase: createCategoryUC,
		FindCategoriesUseCase: findCategoryUC,
	}
}

func (controller *CategoryController) FindCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	output, err := controller.FindCategoriesUseCase.Execute()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadGateway)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadGateway, []error{err}))
		return
	}

	fmt.Println(output)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, output)
}

func (controller *CategoryController) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
		return
	}

	defer r.Body.Close()

	input := dtos.CreateCategoryInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
	}

	output, err := controller.CreateCategoryUseCase.Execute(input)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, output)
}
