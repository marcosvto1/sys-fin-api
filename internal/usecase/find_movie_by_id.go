package usecase

import (
	"errors"

	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg"
)

type FindMovieByIdUC struct {
	movieRepository database.IMovieRepository
}

func NewFindMovieByIdUC(movieRepository database.IMovieRepository) *FindMovieByIdUC {
	return &FindMovieByIdUC{
		movieRepository: movieRepository,
	}
}

func (uc *FindMovieByIdUC) Execute(id int) (dtos.MovieOutput, error) {
	movie, err := uc.movieRepository.FindById(id)
	if err != nil {
		return dtos.MovieOutput{}, err
	}

	if movie.ID == 0 {
		return dtos.MovieOutput{}, errors.New(pkg.ErrNotFoundMovie)
	}

	return dtos.MovieOutput{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Watched:     movie.Watched,
		ImageUrl:    movie.ImageUrl,
	}, nil
}
