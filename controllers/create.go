package controllers

import (
	"encoding/json"
	"fmt"
	"go-test-supporting-project/models"
	"io/ioutil"
	"net/http"
)

type PostCreateRequestObj struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostCreateResponseObj struct {
	Id int `json:"id"`
	ErrorResponseObj
}

type ErrorResponseObj struct {
	Error string `json:"error,omitempty"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var postCreateReq PostCreateRequestObj
	json.Unmarshal(reqBody, &postCreateReq)

	post, err := models.NewPostWithTitle(postCreateReq.Title)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resBody := ErrorResponseObj{
			Error: fmt.Sprintf("Invalid post title: %s", err),
		}
		json.NewEncoder(w).Encode(resBody)
		return
	}

	postCreated, err := h.Store.CreatePost(post)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resBody := ErrorResponseObj{
			Error: fmt.Sprintf("post create failed: %s", err),
		}
		json.NewEncoder(w).Encode(resBody)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resBody := PostCreateResponseObj{
		Id: postCreated.Id,
	}
	json.NewEncoder(w).Encode(resBody)
}
