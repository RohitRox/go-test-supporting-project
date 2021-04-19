package controllers

import (
	"go-test-supporting-project/models"
)

type Handler struct {
	Store StoreIface
}

type StoreIface interface {
	CreatePost(models.Post) (post models.Post, err error)
}
