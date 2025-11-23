package database

import (
	"database/sql"
	"fmt"

	"gitlab.com/marcosvto/sys-fin-api/internal/entity"
)

type FindMovieOptions struct {
	Title   string
	Watched string
}

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) Create(movie *entity.Movie) (*entity.Movie, error) {
	query := "INSERT INTO movies (title, watched, description, image_url) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, movie.Title, movie.Watched, movie.Description, movie.ImageUrl).Scan(&movie.ID)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (r *MovieRepository) FindAll() ([]entity.Movie, error) {
	query := "SELECT id, title, watched, description, image_url FROM movies"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []entity.Movie

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var movie entity.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Watched, &movie.Description, &movie.ImageUrl)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *MovieRepository) Find(offset, pageSize int, filter FindMovieOptions) ([]entity.Movie, int, error) {
	query := "SELECT id, title, watched, description, image_url FROM movies"

	var (
		whereClauses []string
		args         []interface{}
		argPos       = 1
	)

	if filter.Title != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("title ILIKE $%d", argPos))
		args = append(args, "%"+filter.Title+"%")
		argPos++
	}

	print("filter.Watched", filter.Watched)

	if filter.Watched != "" {

		// Convert the string to a boolean
		watched := false
		if filter.Watched == "S" {
			watched = true
		}

		whereClauses = append(whereClauses, fmt.Sprintf("watched = $%d", argPos))
		args = append(args, watched)
		argPos++
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + joinStrings(whereClauses, " AND ")
	}

	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, pageSize, offset)

	fmt.Printf("Query: %s\n", query)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var movies []entity.Movie

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var movie entity.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Watched, &movie.Description, &movie.ImageUrl)
		if err != nil {
			return nil, 0, err
		}
		movies = append(movies, movie)
	}

	return movies, len(movies), nil
}

func (r *MovieRepository) FindById(id int) (entity.Movie, error) {
	query := "SELECT id, title, watched, description, image_url FROM movies WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var movie entity.Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Watched, &movie.Description, &movie.ImageUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Movie{}, nil // No movie found
		}
		return entity.Movie{}, err
	}

	return movie, nil
}

func (r *MovieRepository) Update(movie *entity.Movie) error {
	query := "UPDATE movies SET title = $1, watched = $2, description = $3, image_url= $4 WHERE id = $5"
	_, err := r.db.Exec(query, movie.Title, movie.Watched, movie.Description, movie.ImageUrl, movie.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MovieRepository) DeleteById(id int) error {
	query := "DELETE FROM movies WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *MovieRepository) FindByTitle(title string) (entity.Movie, error) {
	query := "SELECT id, title, watched, description, image_url FROM movies WHERE title = $1"

	row := r.db.QueryRow(query, title)

	var movie entity.Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Watched, &movie.Description, &movie.ImageUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Movie{}, nil // No movie found
		}
		return entity.Movie{}, err
	}

	return movie, nil
}

func joinStrings(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
