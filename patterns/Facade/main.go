// Паттерн Фасад, предоставляет простой интерфейс
// к более сложной системе классов
// Паттерн помогает скрыть детали реализации сложной системы
// и предоставить простой интерфейс для общения со сложной системой
package main

import (
	"facade/facade"
	"facade/user"
)

func main() {
	user := &user.User{
		Name:     "Artem",
		LastName: "Artem",
	}

	userFacade := facade.NewUserFacade() // создаем объект Фасада и взаимодействуем только с ним

	userFacade.CreateUser(user)

	postId := userFacade.CreatePost("testPost")

	userFacade.CreateComment(postId, "testComment")
}
