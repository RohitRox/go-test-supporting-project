package models

import (
	"errors"
	"regexp"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var validPostTitlePatt = regexp.MustCompile(`^\w+[\w\s]+$`)

func NewPostWithTitle(title string) (post Post, err error) {
	if !validPostTitlePatt.MatchString(title) {
		err = errors.New("title is required and only alpha-numeric characters and underscore are permitted in title")
	}
	post = Post{Title: title}

	return
}
