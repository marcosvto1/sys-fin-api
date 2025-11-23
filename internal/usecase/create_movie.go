package usecase

import (
	"errors"
	"log"

	"github.com/lib/pq"
	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg"
)

type CreateMovieUC struct {
	movieRepository database.IMovieRepository
}

func NewCreateMovieUC(movieRepository database.IMovieRepository) *CreateMovieUC {
	return &CreateMovieUC{
		movieRepository: movieRepository,
	}
}

func (uc *CreateMovieUC) Execute(createMovieDto dtos.CreateMovieInput) (dtos.MovieOutput, error) {
	movie := entity.NewMovie(
		-1,
		createMovieDto.Title,
		createMovieDto.Description,
		createMovieDto.ImageUrl,
		createMovieDto.Watched,
	)

	entity, err := uc.movieRepository.Create(movie)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			log.Println("Movie already exists:", err.Error())
			return dtos.MovieOutput{}, errors.New(pkg.ErrAlreadyExistsMovie)
		}

		log.Println("Error creating movie:", err.Error())
		return dtos.MovieOutput{}, err
	}

	return dtos.MovieOutput{
		ID:          entity.ID,
		Title:       entity.Title,
		Watched:     entity.Watched,
		Description: entity.Description,
		ImageUrl:    entity.ImageUrl,
	}, nil
}
