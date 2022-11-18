package utils

import (
	"avatar.com/avatar/avatar/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserIdFromHeader(c *gin.Context) uint64 {
	var userIdStr string = c.Request.Header.Get(config.UserIdHeader)

	userIdUint, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return 0
	}

	return userIdUint
}
