package enteties

type Film struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Ganre       string `json:"ganre"`
	DirectorID  string `json:"director_id"`
	Rate        int    `json:"rate"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
}

type FavoriteList struct {
	ID       string `json:"id"`
	FilmID   string `json:"film_id"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

type WhishList struct {
	ID       string `json:"id"`
	FilmID   string `json:"film_id"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}
