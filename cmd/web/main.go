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
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/web/handlers"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase"

	"github.com/go-chi/jwtauth"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	path, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/.env", path))
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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
		MaxAge:         300,
	}))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	tokenAuth := jwtauth.New("HS256", []byte("any_secret"), nil)
	tokenExpiration := 300
	router.Use(middleware.WithValue("jwt", tokenAuth))
	userRepository := database.NewUserRepository(db)
	createUserUC := usecase.NewCreateUser(userRepository)
	findUserUC := usecase.NewFindUserUseCase(userRepository)
	loginUC := usecase.NewLoginUseCase(userRepository)
	userController := handlers.NewUserController(createUserUC, findUserUC, loginUC, tokenExpiration)

	walletRepository := database.NewWalletRepository(db)
	createWalletUC := usecase.NewCreateWalletUseCase(walletRepository, userRepository)
	walletController := handlers.NewWalletController(createWalletUC)

	categoryRepository := database.NewCategoryRepository(db)
	createCategoryUC := usecase.NewCreateCategoryUseCase(categoryRepository)
	categoryController := handlers.NewCategoryController(createCategoryUC)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/users", func(r chi.Router) {
			r.Post("/", userController.CreateUserHandler)
			r.Get("/", userController.FindUserHandler)
		})

		r.Route("/wallets", func(r chi.Router) {
			r.Post("/", walletController.CreateWalletHandler)
		})

		r.Route("/categories", func(r chi.Router) {
			r.Post("/", categoryController.CreateCategoryHandler)
		})
	})

	router.Post("/auth/access_token", userController.LoginHandler)

	err = http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
