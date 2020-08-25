package user

import (
	"github.com/gin-gonic/gin"

	. "gin-api-boilerplate/handler"
	"gin-api-boilerplate/model"
	"gin-api-boilerplate/pkg/auth"
	"gin-api-boilerplate/pkg/ecode"
	"gin-api-boilerplate/pkg/token"
)

func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, ecode.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, ecode.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, ecode.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, ecode.ErrToken, nil)
		return
	}

	SendResponse(c, model.Token{Token: t}, nil)
}
