package teamAvatars

import (
	"avatar.com/avatar/avatar/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetTeamAvatarPath Get src-path to user's avatar
func GetTeamAvatarPath(c *gin.Context) {
	teamIdStr := c.Param(TeamIdParamName)

	teamId, err := strconv.ParseUint(teamIdStr, 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Incorrect type of `%s` (should be uint64)", TeamIdParamName)})
		return
	}

	srcPath, err := db.GetTeamSrcPath(teamId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"detail": "Image not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"avatarSrc": srcPath})
}
