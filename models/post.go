package models

// Post type details
type Post struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}

type Album struct {
	ID     int64  `json:"id"`
	UserID string `json:"userId"`
	Title  string `json:"title"`
}

type Photo struct {
	ID           int64  `json:"id"`
	AlbumID      string `json:"albumId"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type SearchRes struct {
	UserID       string `json:"userId"`
	AlbumID      string `json:"albumId"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
