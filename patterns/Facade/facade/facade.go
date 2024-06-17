package facade

import (
	"facade/comment"
	"facade/post"
	"facade/user"
)

// Аккумулируем всю сложную логику в одном месте и отдаем только методы для работы с ней
type UserFacade struct {
	user    *user.User
	post    *post.Post
	comment *comment.Comment
}

func NewUserFacade() *UserFacade {
	return &UserFacade{
		user:    &user.User{},
		post:    &post.Post{},
		comment: &comment.Comment{},
	}
}

func (uf *UserFacade) CreateUser(u *user.User) {
	uf.user.NewUser(u.Name, u.LastName)
}

func (uf *UserFacade) CreatePost(text string) int {
	post := uf.post.NewPost(text)
	return post.Id
}

func (uf *UserFacade) CreateComment(postId int, text string) {
	uf.comment.NewComment(postId, text)
}
