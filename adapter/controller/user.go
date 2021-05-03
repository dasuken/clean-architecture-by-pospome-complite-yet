package controller

import (
	"fmt"
	"github.com/noguchidaisuke/clean-architecture/usecase"
	"net/http"
)

type User struct {
	InputFactory func(o usecase.UserOutputPort) usecase.UserInputPort
	OutputFactory func(w http.ResponseWriter) usecase.UserOutputPort
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := 1

	outputPort := u.OutputFactory(w)
	inputPort := u.InputFactory(outputPort)
	err := inputPort.GetUserByID(userID)
	if err != nil {
		fmt.Fprint(w, "error")
	}
}
