package utils

import (
	"avatar.com/avatar/server/conf"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserIdFromHeader(c *gin.Context) uint64 {
	var userIdStr string = c.Request.Header.Get(conf.UserIdHeader)

	userIdUint, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return 0
	}

	return userIdUint
}