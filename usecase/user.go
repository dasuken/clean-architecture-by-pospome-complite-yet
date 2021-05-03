package usecase

import "github.com/noguchidaisuke/clean-architecture/entity"

type UserInputPort interface {
	GetUserByID(userID int) error
}

type UserOutputPort interface {
	Render(*entity.User)
}

type User struct {
	output UserOutputPort
}

func NewUser() UserInputPort {
	return &User{}
}

func (u *User) GetUserByID(userID int) error {
	user := &entity.User{
		Id:   userID,
		Name: "JOKER",
	}

	u.output.Render(user)

	return nil
}

