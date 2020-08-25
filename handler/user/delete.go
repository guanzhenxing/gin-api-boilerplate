package user

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "gin-api-boilerplate/handler"
	"gin-api-boilerplate/model"
	"gin-api-boilerplate/pkg/ecode"
)

// @Summary Delete an user by the user identifier
// @Description Delete user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	//直接调用model的方法
	if err := model.DeleteUser(uint64(userId)); err != nil {
		Error(c, ecode.ErrDatabase)
		return
	}

	Success(c, nil)
}
