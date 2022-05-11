package handlers

import (
	"avatar.com/avatar/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetAvatarPath Get src-path to user's avatar
func GetAvatarPath(c *gin.Context) {
	userIdStr := c.Request.URL.Query().Get("user")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect type of `user` (should be uint64)"})
		return
	}

	srcPath, err := db.GetSrcPath(userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"avatarSrc": srcPath})
}
