package entity

type Movie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Watched     bool   `json:"watched"`
	ImageUrl    string `json:"image_url"`
}

func NewMovie(id int64, title, description, image_url string, watched bool) *Movie {
	return &Movie{
		ID:          id,
		Title:       title,
		Watched:     watched,
		Description: description,
		ImageUrl:    image_url,
	}
}

func (m *Movie) MarkAsWatched() {
	m.Watched = true
}

func (m *Movie) MarkAsUnwatched() {
	m.Watched = false
}

func (m *Movie) IsWatched() bool {
	return m.Watched
}

func (m *Movie) GetID() int64 {
	return m.ID
}

func (m *Movie) GetTitle() string {
	return m.Title
}

func (m *Movie) SetTitle(title string) {
	m.Title = title
}
func (m *Movie) GetWatched() bool {
	return m.Watched
}
func (m *Movie) SetWatched(watched bool) {
	m.Watched = watched
}
func (m *Movie) String() string {
	return m.Title
}
func (m *Movie) ToString() string {
	return m.Title
}

func (m *Movie) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":      m.ID,
		"title":   m.Title,
		"watched": m.Watched,
	}
}
