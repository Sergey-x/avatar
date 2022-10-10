package handlers

import (
	"avatar.com/avatar/db"
	"avatar.com/avatar/server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetAvatarPath Get src-path to user's avatar
func GetAvatarPath(c *gin.Context) {
	var userId uint64 = utils.GetUserIdFromHeader(c)
	if userId == 0 {
		log.Println("User id = 0")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect type of `user id` (should be uint64)"})
		return
	}

	srcPath, err := db.GetSrcPath(userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"detail": "Image not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"avatarSrc": srcPath})
}
