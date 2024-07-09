package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg/errorable"
)

type UserController struct {
	CreateUserUseCase *usecase.CreateUserUseCase
	FindUserUseCase   *usecase.FindUserUseCase
	LoginUseCase      usecase.ILoginUseCase
	JwtExpiresIn      int
}

func NewUserController(
	createUserUC *usecase.CreateUserUseCase,
	findUserUC *usecase.FindUserUseCase,
	loginUseCase usecase.ILoginUseCase,
	jwtExpiresIn int,
) *UserController {
	return &UserController{
		CreateUserUseCase: createUserUC,
		FindUserUseCase:   findUserUC,
		LoginUseCase:      loginUseCase,
		JwtExpiresIn:      jwtExpiresIn,
	}
}

func (this *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	input := &dtos.CreateUserInput{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.NewHttpError(http.StatusBadRequest, []errorable.CtxError{
			*errorable.New(errorable.INVALID_BODY_REQUEST),
		}))
		return
	}

	err = json.Unmarshal(body, input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.NewHttpError(http.StatusBadRequest, []errorable.CtxError{
			*errorable.New(errorable.INVALID_BODY_REQUEST),
		}))
		return
	}

	output, err := this.CreateUserUseCase.Execute(*input)
	if err != nil {
		render.JSON(w, r, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, output)
}

func (this *UserController) FindUserHandler(w http.ResponseWriter, r *http.Request) {
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

	input := dtos.FindInput{
		Id:         id,
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}

	output, err := this.FindUserUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, errorable.NewHttpError(http.StatusInternalServerError, []errorable.CtxError{
			*errorable.New(errorable.INTERNAL_ERROR),
		}))
		return
	}

	json.NewEncoder(w).Encode(output)
}

func (this *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var errors []map[string]string
	input := dtos.LoginInput{}
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errors = append(errors, map[string]string{
			"field":   "body",
			"message": err.Error(),
		})
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &input)

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	output, error := this.LoginUseCase.Execute(jwt, this.JwtExpiresIn, input)
	if error != nil {
		var status int
		if error.Context == errorable.INVALID_PASSWORD {
			status = http.StatusBadRequest
		} else if error.Context == errorable.NOT_FOUND_REGISTER {
			status = http.StatusNotFound
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errorable.NewHttpError(status, []errorable.CtxError{
			*error,
		}))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
