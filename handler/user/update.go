package user

import (
	"gin-api-boilerplate/pkg/log"
	"go.uber.org/zap"
	"strconv"

	. "gin-api-boilerplate/handler"
	"gin-api-boilerplate/model"
	"gin-api-boilerplate/pkg/ecode"
	"gin-api-boilerplate/util"

	"github.com/gin-gonic/gin"
)

// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Param user body model.UserModel true "The user info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [put]
func Update(c *gin.Context) {
	log.Info("Update function called.", zap.Any("X-Request-Id", util.GetReqID(c)))

	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		Error(c, ecode.ErrBind)
		return
	}

	// We update the record based on the user id.
	u.Id = uint64(userId)

	// Validate the data.
	if err := u.Validate(); err != nil {
		Error(c, ecode.ErrValidation)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		Error(c, ecode.ErrEncrypt)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		Error(c, ecode.ErrDatabase)
		return
	}

	Success(c, nil)
}
