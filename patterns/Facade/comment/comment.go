package comment

import (
	"fmt"
	"math/rand"
)

type Comment struct {
	id     int
	postId int
	text   string
}

func (p *Comment) NewComment(postId int, text string) *Comment {
	id := rand.Int()
	fmt.Printf("comment with id: %d and text: %s has been created \n", id, text)
	return &Comment{
		id:     id,
		postId: postId,
		text:   text,
	}
}

// more code...
