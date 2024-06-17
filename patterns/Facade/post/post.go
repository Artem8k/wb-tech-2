package post

import (
	"fmt"
	"math/rand"
)

type Post struct {
	Id   int
	text string
}

func (p *Post) NewPost(text string) *Post {
	id := rand.Int()
	fmt.Printf("post with id: %d and text: %s has been created \n", id, text)
	return &Post{
		Id:   id,
		text: text,
	}
}

// more code...
