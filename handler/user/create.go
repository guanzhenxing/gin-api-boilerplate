package user

import (
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	. "gin-api-boilerplate/handler"
	"gin-api-boilerplate/model"
	"gin-api-boilerplate/pkg/ecode"
	"gin-api-boilerplate/pkg/log"
	"gin-api-boilerplate/util"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	log.Info("User Create function called.", zap.Any("X-Request-Id", util.GetReqID(c)))

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		Error(c, ecode.ErrBind)
		return
	}

	//这里直接从controller调用到model
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, ecode.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, ecode.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, ecode.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	Success(c, rsp)
}
