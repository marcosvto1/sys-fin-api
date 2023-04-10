package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-adv-api/pkg/errorable"
)

type CategoryController struct {
	CreateCategoryUseCase usecase.ICreateCategory
}

func NewCategoryController(createCategoryUC usecase.ICreateCategory) *CategoryController {
	return &CategoryController{
		CreateCategoryUseCase: createCategoryUC,
	}
}

func (this *CategoryController) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
	}

	defer r.Body.Close()

	input := dtos.CreateCategoryInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
	}

	output, err := this.CreateCategoryUseCase.Execute(input)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadRequest, []error{err}))
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, output)
}
