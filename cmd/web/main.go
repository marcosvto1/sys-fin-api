package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/web/handlers"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"

	"github.com/go-chi/jwtauth"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	path, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/.env", path))
	if err != nil {
		log.Warn("Error loading environment variables file")
	}

	fmt.Println(os.Getenv("DATABASE_URL"))

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/health"))
	router.Use(middleware.AllowContentType("application/json", "text/xml"))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         30000,
	}))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	tokenAuth := jwtauth.New("HS256", []byte("any_secret"), nil)
	tokenExpiration := 30000
	router.Use(middleware.WithValue("jwt", tokenAuth))
	userRepository := database.NewUserRepository(db)
	createUserUC := usecase.NewCreateUser(userRepository)
	findUserUC := usecase.NewFindUserUseCase(userRepository)
	loginUC := usecase.NewLoginUseCase(userRepository)
	userController := handlers.NewUserController(createUserUC, findUserUC, loginUC, tokenExpiration)

	categoryRepository := database.NewCategoryRepository(db)
	createCategoryUC := usecase.NewCreateCategoryUseCase(categoryRepository)
	findCategoryUC := usecase.NewFindCategoriesUseCase(categoryRepository)
	categoryController := handlers.NewCategoryController(createCategoryUC, findCategoryUC)

	walletRepository := database.NewWalletRepository(db)
	createWalletUC := usecase.NewCreateWalletUseCase(walletRepository, userRepository)
	findWalletsUC := usecase.NewFindWalletsUseCase(walletRepository)
	walletController := handlers.NewWalletController(createWalletUC, findWalletsUC)

	transactionRepository := database.NewTransactionRepository(db)
	createTransactionUC := usecase.NewCreateTransactionUseCase(transactionRepository)
	findTransactionUC := usecase.NewFindTransactionUseCase(transactionRepository)
	findOneTransactionUC := usecase.NewFindOneTransactionUseCase(transactionRepository)
	deleteTransactinoUC := usecase.NewDeleteTransactionUsecase(transactionRepository)
	updateTransactinoUC := usecase.NewUpdateTransactoinUseCase(transactionRepository)
	transactionController := handlers.NewTransactionController(
		createTransactionUC,
		findTransactionUC,
		findOneTransactionUC,
		deleteTransactinoUC,
		updateTransactinoUC,
	)

	getChartTransactionByCategoryUC := usecase.NewChartTransactionByCategoryUseCase(transactionRepository)
	getChartTransactionByTypeUC := usecase.NewChartTransactionByTypeUseCase(transactionRepository)
	dashboardController := handlers.NewDashboardController(getChartTransactionByCategoryUC, getChartTransactionByTypeUC)

	// SUBSCRIPTION
	subscriptionRepository := database.NewSubscriptionRepository(db)
	findSubscriptionUC := usecase.NewFindSubscriptionUsecase(*subscriptionRepository)
	createSubscriptionUC := usecase.NewCreateSubscriptionUsecase(*subscriptionRepository)
	deleteSubscriptionUC := usecase.NewDeleteSubscriptionUsecase(subscriptionRepository)
	updateSubscriptionUC := usecase.NewUpdateSubscriptionUsecase(subscriptionRepository)

	subscriptionHandler := handlers.NewSubscriptionHandler(
		findSubscriptionUC,
		createSubscriptionUC,
		deleteSubscriptionUC,
		updateSubscriptionUC,
	)
	// -- END SUBSCRIPTION

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		// ro

		r.Route("/users", func(r chi.Router) {
			r.Get("/", userController.FindUserHandler)
		})

		r.Route("/wallets", func(r chi.Router) {
			r.Post("/", walletController.CreateWalletHandler)
			r.Get("/", walletController.FindWalletsHandler)
		})

		r.Route("/categories", func(r chi.Router) {
			r.Post("/", categoryController.CreateCategoryHandler)
			r.Get("/", categoryController.FindCategoriesHandler)
		})

		r.Route("/transactions", func(r chi.Router) {
			r.Get("/", transactionController.FindTransactionHandler)
			r.Get("/{id}", transactionController.FindOneTransactionHandler)
			r.Post("/", transactionController.CreateTransactionHandler)
			r.Put("/{id}", transactionController.UpdateTransactionHandler)
			r.Delete("/{id}", transactionController.DeleteTransactionHandler)
		})

		r.Route("/dashboard", func(r chi.Router) {
			r.Get("/by-category", dashboardController.GetChartTransactionByCategoryHandler)
			r.Get("/by-type", dashboardController.GetChartTransactionByTypeHandler)
		})

		r.Route("/subscriptions", func(r chi.Router) {
			r.Get("/", subscriptionHandler.FindAllHandler)
			r.Post("/", subscriptionHandler.CreateHandler)
			r.Get("/{id}", subscriptionHandler.FindByIdHandler)
			r.Delete("/{id}", subscriptionHandler.DeleteByIdHandler)
			r.Put("/{id}", subscriptionHandler.UpdateByIdHandler)
		})
	})

	router.Post("/auth/access_token", userController.LoginHandler)
	router.Post("/users", userController.CreateUserHandler)

	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
	if err != nil {
		log.Fatal(err)
	}
}
