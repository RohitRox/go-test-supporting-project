package models_test

import (
	"testing"

	m "go-test-supporting-project/models"
)

func TestPost(t *testing.T) {
	testData := []struct {
		title string
		error bool
	}{
		{"Hello World", false},
		{"Hello Testing 124", false},
		{"Hello_World", false},
		{"Hello World!", true},
		{"Hello World - 124", true},
		{"Hello@World", true},
	}

	for _, dat := range testData {
		t.Run(dat.title, func(t *testing.T) {
			post, err := m.NewPostWithTitle(dat.title)
			if dat.error {
				if err == nil {
					t.Errorf("Expected error Got nil for post: %s", post.Title)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %s for post: %s", err, post.Title)
				}
			}
		})
	}
}
