package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-adv-api/pkg/errorable"

	log "github.com/sirupsen/logrus"
)

type WalletController struct {
	CreateWalletUseCase usecase.ICreateWallet
	FindWalletsUsecase  usecase.IFindWallets
}

func NewWalletController(walletUseCase usecase.ICreateWallet, findWalletsUseCase usecase.IFindWallets) *WalletController {
	return &WalletController{
		CreateWalletUseCase: walletUseCase,
		FindWalletsUsecase:  findWalletsUseCase,
	}
}

func (controller *WalletController) FindWalletsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := controller.FindWalletsUsecase.Execute()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadGateway)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadGateway, []error{err}))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, output)
}

func (controller *WalletController) CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.NewHttpError(http.StatusBadRequest, []errorable.CtxError{
			*errorable.New(errorable.INVALID_BODY_REQUEST),
		}))
		return
	}
	defer r.Body.Close()

	input := dtos.CreateWalletInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorable.NewHttpError(http.StatusBadRequest, []errorable.CtxError{
			*errorable.New(errorable.INVALID_BODY_REQUEST),
		}))
		return
	}

	output, err := controller.CreateWalletUseCase.Execute(input)
	log.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		render.JSON(w, r, errorable.HttpResponse(http.StatusBadGateway, []error{err}))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, output)
}
