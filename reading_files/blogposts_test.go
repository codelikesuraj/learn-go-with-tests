package reading_files

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello_world2.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
