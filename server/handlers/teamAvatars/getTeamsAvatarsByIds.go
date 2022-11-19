package teamAvatars

import (
	"avatar.com/avatar/db"
	"avatar.com/avatar/server/schemas"
	"avatar.com/avatar/server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetTeamsAvatarsByIds Get avatar for teams
func GetTeamsAvatarsByIds(c *gin.Context) {

	var teamsIdsParam schemas.TeamsIdsParam

	err := c.ShouldBindQuery(&teamsIdsParam)
	if err != nil {
		log.Println("Bad teams ids")
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Bad teams ids format"})
		return
	}
	teamsIds := utils.StrToArrayInt(teamsIdsParam.TeamsIds)
	if len(teamsIds) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "No correct ids"})
		return
	}

	avatarsPaths := db.GetTeamsAvatarsByIds(teamsIds)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"detail": "Image not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"avatars": avatarsPaths})
	return
}
