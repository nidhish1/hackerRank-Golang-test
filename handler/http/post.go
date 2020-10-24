package handler

import (
	"encoding/json"
	"net/http"

	"hackerRank-Golang-test/driver"
	repository "hackerRank-Golang-test/repository"
	post "hackerRank-Golang-test/repository/post"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

// Post ...
type Post struct {
	repo repository.PostRepo
}

func (p *Post) Search(w http.ResponseWriter, r *http.Request) {
	//uid, _ := strconv.Atoi(chi.URLParam(r, "id"))
	//albumId, _ := strconv.Atoi(chi.URLParam(r, "aid"))

	payload, err := p.repo.Fetch(r.Context(), 1, 1)

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

//respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
