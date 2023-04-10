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
}

func NewWalletController(walletUseCase usecase.ICreateWallet) *WalletController {
	return &WalletController{
		CreateWalletUseCase: walletUseCase,
	}
}

func (this *WalletController) CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
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

	output, err := this.CreateWalletUseCase.Execute(input)
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
