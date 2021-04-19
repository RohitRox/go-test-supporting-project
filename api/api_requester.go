package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	m "go-test-supporting-project/models"
)

type ApiRequester struct {
	BaseUrl string
}

func (a ApiRequester) GetPost(id int) (post m.Post, err error) {
	url := fmt.Sprintf(a.BaseUrl+"/posts/%d", id)

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&post)

	if err != nil {
		return
	}

	return
}

func (a ApiRequester) GetPosts() (posts []m.Post, err error) {
	url := fmt.Sprintf(a.BaseUrl + "/posts")

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&posts)

	if err != nil {
		return
	}

	return
}

func (a ApiRequester) CreatePost(post m.Post) (postCreated m.Post, err error) {
	url := fmt.Sprintf(a.BaseUrl + "/posts")

	postBody, err := json.Marshal(post)

	if err != nil {
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(postBody))

	if err != nil {
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&postCreated)

	return
}
