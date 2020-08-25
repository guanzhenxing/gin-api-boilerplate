package user

import (
	"github.com/gin-gonic/gin"

	. "gin-api-boilerplate/handler"
	"gin-api-boilerplate/model"
	"gin-api-boilerplate/pkg/ecode"
)

// @Summary Get an user by the user identifier
// @Description Get an user by username
// @Tags user
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.UserModel "{"code":0,"message":"OK","data":{}}"
// @Router /user/{username} [get]
func Get(c *gin.Context) {

	username := c.Param("username")

	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		Error(c, ecode.ErrUserNotFound)
		return
	}

	Success(c, user)
}
