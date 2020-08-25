package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"

	"gin-api-boilerplate/config/constvar"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

func IsDebugMode(env string) bool {
	return InArray(env, []string{constvar.EnvSOLO, constvar.EnvDEV})
}

func IsTestMode(env string) bool {
	return InArray(env, []string{constvar.EnvQA})
}

func IsReleaseMode(env string) bool {
	return InArray(env, []string{constvar.EnvUAT, constvar.EnvPROD})
}
