package api_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	a "go-test-supporting-project/api"
	"go-test-supporting-project/models"
)

func TestApiRequester(t *testing.T) {
	postID := 2101
	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/posts":
				switch r.Method {
				case "POST":
					body, _ := ioutil.ReadAll(r.Body)
					var post models.Post
					err := json.Unmarshal(body, &post)

					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintf(w, "Error")
					}

					post.Id = postID
					json.NewEncoder(w).Encode(post)
					return
				}
			}
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid route")
		}),
	)

	defer testServer.Close()

	t.Run("CreatePost", func(t *testing.T) {
		apiRequester := a.ApiRequester{
			BaseUrl: testServer.URL,
		}

		post := models.Post{Title: "Hello 124"}

		postCreated, err := apiRequester.CreatePost(post)

		if err != nil {
			t.Fatal(err)
		}

		if postCreated.Id != postID {
			t.Fatalf("Expected post id: %d got: %d", postID, postCreated.Id)
		}
	})
}
