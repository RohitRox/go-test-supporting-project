package db

import "go-test-supporting-project/models"

type Storage struct {
}

func (s Storage) CreatePost(models.Post) (post models.Post, err error) {
	return
}
