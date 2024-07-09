package webrouters

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/web/handlers"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"
)

func CategorySetupRouter(db *sql.DB, r chi.Router) {
	categoryRepository := database.NewCategoryRepository(db)
	createCategoryUC := usecase.NewCreateCategoryUseCase(categoryRepository)
	findCategoryUC := usecase.NewFindCategoriesUseCase(categoryRepository)
	categoryController := handlers.NewCategoryController(createCategoryUC, findCategoryUC)

	r.Route("/categories", func(r chi.Router) {
		r.Post("/", categoryController.CreateCategoryHandler)
		r.Get("/", categoryController.FindCategoriesHandler)
	})
}
