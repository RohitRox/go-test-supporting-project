package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	c "go-test-supporting-project/controllers"
	m "go-test-supporting-project/models"
)

type MockedStore struct {
	HandleCreatePost func(m.Post) (post m.Post, err error)
}

func (m MockedStore) CreatePost(postBody m.Post) (post m.Post, err error) {
	if m.HandleCreatePost != nil {
		return m.HandleCreatePost(postBody)
	}
	return
}

func TestCreateHandler(t *testing.T) {

	t.Log("success usecase")
	{
		testData := []struct {
			title string
		}{
			{"Hello World"},
			{"Hello Testing 124"},
		}

		var idCounter = 1101

		mockedStore := &MockedStore{
			HandleCreatePost: func(postBody m.Post) (post m.Post, err error) {
				post = m.Post{
					Id:    idCounter,
					Title: postBody.Title,
				}
				return
			},
		}

		for _, dat := range testData {
			t.Run("post create", func(t *testing.T) {
				postBody := c.PostCreateRequestObj{Title: dat.title}
				body, err := json.Marshal(postBody)

				if err != nil {
					t.Fatal(err)
				}

				req, _ := http.NewRequest("POST", "/postcreateroute", bytes.NewReader(body))

				if err != nil {
					t.Fatal(err)
				}

				handlers := c.Handler{
					Store: mockedStore,
				}

				rec := httptest.NewRecorder()
				handlers.Create(rec, req)

				status := rec.Code

				if status != 201 {
					t.Errorf("Expected post create with title: %s to success, failed with status: %d", dat.title, status)
					return
				}

				var resJson map[string]int
				_ = json.Unmarshal(rec.Body.Bytes(), &resJson)

				if resJson["id"] != idCounter {
					t.Errorf("Expected id of post created: %v got: %v", idCounter, resJson["id"])
				}
			})
		}
	}

	t.Log("validation failure usecase")
	{
		testData := []struct {
			title string
		}{
			{"Hello-World"},
			{"Hello! 124"},
		}

		mockedStore := &MockedStore{}

		for _, dat := range testData {
			t.Run("post create", func(t *testing.T) {
				postBody := c.PostCreateRequestObj{Title: dat.title}
				body, err := json.Marshal(postBody)

				if err != nil {
					t.Fatal(err)
				}

				req, _ := http.NewRequest("POST", "/postcreateroute", bytes.NewReader(body))

				if err != nil {
					t.Fatal(err)
				}

				handlers := c.Handler{
					Store: mockedStore,
				}

				rec := httptest.NewRecorder()
				handlers.Create(rec, req)

				status := rec.Code

				if status != 400 {
					t.Errorf("Expected post request to dail with status: %d got: %d", 400, status)
					return
				}

				var resJson map[string]string
				_ = json.Unmarshal(rec.Body.Bytes(), &resJson)

				if len(resJson["error"]) < 1 {
					t.Error("Expected error message to be present, got empty")
				}
			})
		}
	}

	t.Log("external api request failure usecase")
	{
		testData := struct {
			title string
		}{
			"Hello Testing 124",
		}

		mockedStore := &MockedStore{
			HandleCreatePost: func(postBody m.Post) (post m.Post, err error) {
				err = fmt.Errorf("Network error")
				return
			},
		}

		t.Run("post create", func(t *testing.T) {
			postBody := c.PostCreateRequestObj{Title: testData.title}
			body, err := json.Marshal(postBody)

			if err != nil {
				t.Fatal(err)
			}

			req, _ := http.NewRequest("POST", "/postcreateroute", bytes.NewReader(body))

			if err != nil {
				t.Fatal(err)
			}

			handlers := c.Handler{
				Store: mockedStore,
			}

			rec := httptest.NewRecorder()
			handlers.Create(rec, req)

			status := rec.Code
			if status != 400 {
				t.Errorf("Expected post request to dail with status: %d got: %d", 400, status)
				return
			}

			var resJson map[string]string
			_ = json.Unmarshal(rec.Body.Bytes(), &resJson)

			if len(resJson["error"]) < 1 {
				t.Error("Expected error message to be present, got empty")
			}
		})
	}
}
