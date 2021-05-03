package presenter

import (
	"fmt"
	"github.com/noguchidaisuke/clean-architecture/entity"
	"github.com/noguchidaisuke/clean-architecture/usecase"
	"net/http"
)

type User struct {
	w http.ResponseWriter
}

func NewUser(w http.ResponseWriter) usecase.UserOutputPort {
	return &User{
		w: w,
	}
}

func (u *User) Render(user *entity.User) {
	fmt.Fprintln(u.w, user.Name)
}