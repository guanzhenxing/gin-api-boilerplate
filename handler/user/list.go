package user

import (
	"github.com/gin-gonic/gin"

	. "gin-api-boilerplate/handler"
	"gin-api-boilerplate/model"
	"gin-api-boilerplate/pkg/ecode"
	"gin-api-boilerplate/pkg/log"
	"gin-api-boilerplate/service"
)

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

// @Summary List the users in the database
// @Description List users
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.ListRequest true "List users"
// @Success 200 {object} user.ListResponse "{"code":0,"message":"OK","data":{}}"
// @Router /user [get]
func List(c *gin.Context) {

	log.Info("List function called.")

	var r ListRequest
	if err := c.Bind(&r); err != nil {
		Error(c, ecode.ErrBind)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		Error(c, err)
		return
	}

	Success(c, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
