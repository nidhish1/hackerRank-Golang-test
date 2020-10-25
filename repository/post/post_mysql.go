package post

import (
	"context"
	"database/sql"

	"hackerRank-Golang-test/models"
	pRepo "hackerRank-Golang-test/repository"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewSQLPostRepo(Conn *sql.DB) pRepo.PostRepo {
	return &mysqlPostRepo{
		Conn: Conn,
	}
}

type mysqlPostRepo struct {
	Conn *sql.DB
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.SearchRes, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.SearchRes, 10)
	for rows.Next() {
		data := new(models.SearchRes)

		err := rows.Scan(
			&data.UserID,
			&data.Title,
			&data.AlbumID,
			&data.ThumbnailURL,
			&data.URL,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlPostRepo) Fetch(ctx context.Context, uid int, albumId int) ([]*models.SearchRes, error) {
	query := "SELECT  album.userId, album.title,photo.albumId, photo.thumbnailUrl ,photo.url from photo join album on photo.id = album.id where album.userId = ? and photo.albumId= ?"

	return m.fetch(ctx, query, uid, albumId)
}
