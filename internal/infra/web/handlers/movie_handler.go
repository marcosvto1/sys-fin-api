package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg"
	"gitlab.com/marcosvto/sys-fin-api/pkg/httputils"
)

type MovieHandler struct {
	CreateMovie   usecase.ICreateMovie
	FindMovie     usecase.IFindMovies
	DeleteMovie   usecase.IDeleteMovie
	FindMovieById usecase.IFindMovieById
	UpdateMovie   usecase.IUpdateMovie
}

func NewMovieHandler(
	createMovie usecase.ICreateMovie,
	findMovie usecase.IFindMovies,
	deleteMovie usecase.IDeleteMovie,
	findMovieById usecase.IFindMovieById,
	updateMovie usecase.IUpdateMovie,
) *MovieHandler {

	return &MovieHandler{
		CreateMovie:   createMovie,
		FindMovie:     findMovie,
		DeleteMovie:   deleteMovie,
		FindMovieById: findMovieById,
		UpdateMovie:   updateMovie,
	}
}

func (h *MovieHandler) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var createMovieDto dtos.CreateMovieInput
	if err := json.NewDecoder(r.Body).Decode(&createMovieDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the use case to create a movie
	movie, err := h.CreateMovie.Execute(createMovieDto)
	if err != nil {
		if err.Error() == pkg.ErrAlreadyExistsMovie {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		// Generic error handling
		http.Error(w, "Failed to create movie", http.StatusInternalServerError)
		return
	}

	// Return the created movie as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) FindMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var findMovieDto dtos.FindMovieInput

	title := r.URL.Query().Get("title")
	watched := r.URL.Query().Get("watched")
	pageNumberStr := r.URL.Query().Get("page_number")
	pageSizeStr := r.URL.Query().Get("page_size")

	// Validate the query parameters

	pageNumber, pageSize, errList := httputils.GetRequestPaginator(pageNumberStr, pageSizeStr)
	if len(errList) > 0 {
		http.Error(w, "Invalid page number or size", http.StatusBadRequest)
		return
	}
	findMovieDto.PageNumber = pageNumber
	findMovieDto.PageSize = pageSize
	if title != "" {
		findMovieDto.Title = title
	}

	if watched != "" {
		if watched != "S" && watched != "N" {
			http.Error(w, "Invalid watched parameter", http.StatusBadRequest)
			return
		}

		findMovieDto.Watched = watched
	}

	moviesPaginator, err := h.FindMovie.Execute(findMovieDto)
	if err != nil {
		log.Printf("Error finding movies: %v", err)
		http.Error(w, "Failed to find movies", http.StatusInternalServerError)
		return
	}

	// // Return the found movies as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moviesPaginator)
}

func (h *MovieHandler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Get the movie ID from the URL
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Movie ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Call the use case to delete the movie
	err = h.DeleteMovie.Execute(id)
	if err != nil {
		http.Error(w, "Failed to delete movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MovieHandler) FindMovieByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Get the movie ID from the URL

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Movie ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Call the use case to find the movie by ID
	movie, err := h.FindMovieById.Execute(id)
	if err != nil {
		log.Printf("Error finding movie by ID: %v", err)
		http.Error(w, "Movie not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Get the movie ID from the URL
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Movie ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Parse the request body
	var updateMovieDto dtos.UpdateMovieInput
	if err := json.NewDecoder(r.Body).Decode(&updateMovieDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println(updateMovieDto)

	// Call the use case to update the movie
	movie, err := h.UpdateMovie.Execute(id, updateMovieDto)
	if err != nil {
		log.Printf("Error updating movie: %v", err)

		if err.Error() == pkg.ErrNotFoundMovie {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err.Error() == pkg.ErrAlreadyExistsMovie {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, "Failed to update movie", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
