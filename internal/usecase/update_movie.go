package usecase

import (
	"errors"
	"log"

	"github.com/lib/pq"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg"
)

type UpdateMovieUC struct {
	movieRepository database.IMovieRepository
}

func NewUpdateMovieUC(movieRepository database.IMovieRepository) *UpdateMovieUC {
	return &UpdateMovieUC{
		movieRepository: movieRepository,
	}
}

func (uc *UpdateMovieUC) Execute(id int, updateMovieDto dtos.UpdateMovieInput) (dtos.MovieOutput, error) {
	movie, err := uc.movieRepository.FindById(id)
	if err != nil {
		log.Println("Error finding movie by ID after update:", err)

		return dtos.MovieOutput{}, err
	}

	log.Println("Movie found:", movie.ID, movie.Title)
	if movie.ID == 0 {
		log.Println("Movie not found")
		return dtos.MovieOutput{}, errors.New(pkg.ErrNotFoundMovie)
	}

	if updateMovieDto.Title != "" {
		movie.Title = updateMovieDto.Title
	}

	movie.Watched = updateMovieDto.Watched

	err = uc.movieRepository.Update(&movie)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			log.Println("Movie already exists:", err.Error())
			return dtos.MovieOutput{}, errors.New(pkg.ErrAlreadyExistsMovie)
		}
		return dtos.MovieOutput{}, err
	}

	entity, err := uc.movieRepository.FindById(id)
	if err != nil {
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
