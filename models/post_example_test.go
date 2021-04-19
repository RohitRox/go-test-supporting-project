package models_test

import (
	"fmt"
	m "go-test-supporting-project/models"
)

func ExamplePost() {
	title := "Hello Testing 124"
	post, err := m.NewPostWithTitle(title)

	if err != nil {
		fmt.Println("Invalid title")
		fmt.Println(err)
	}

	fmt.Printf("Post initialied with title: %s", post.Title)
	// Output:
	// Post initialied with title: Hello Testing 124
}
