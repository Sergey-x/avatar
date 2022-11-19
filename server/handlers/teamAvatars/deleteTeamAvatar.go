package teamAvatars

import (
	"avatar.com/avatar/db"
	"avatar.com/avatar/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// DeleteTeamAvatar Delete user's avatar
func DeleteTeamAvatar(c *gin.Context) {
	// get user id from http header
	var userId uint64 = utils.GetUserIdFromHeader(c)
	if userId == 0 {
		log.Println("User id = 0")
		c.JSON(http.StatusBadRequest, gin.H{"Detail": "Ошибка пользователя - нераспознанный id"})
		return
	}

	teamIdStr := c.Param(TeamIdParamName)
	teamId, err := strconv.ParseUint(teamIdStr, 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Incorrect type of `%s` (should be uint64)", TeamIdParamName)})
		return
	}

	isUserMember := utils.CheckUserIsMemberOfTeam(userId, teamId)
	if isUserMember == false {
		c.JSON(http.StatusForbidden, gin.H{"Detail": "Permission denied"})
		return
	}

	e := db.DeleteTeamSrcPath(teamId)
	if e != nil {
		c.JSON(http.StatusNotFound, gin.H{"Detail": "Аватар не сущетсвует"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"Detail": "Аватар удален"})
}
