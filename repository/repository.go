package repsitory

import (
	"context"

	"hackerRank-Golang-test/models"
)

// PostRepo explain...
type PostRepo interface {
	Fetch(ctx context.Context, uid int, albumId int) ([]*models.SearchRes, error)
}
