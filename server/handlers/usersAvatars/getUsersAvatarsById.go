package usersAvatars

import (
	"avatar.com/avatar/db"
	"avatar.com/avatar/server/schemas"
	"avatar.com/avatar/server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetUsersAvatarsByIds Get avatar for list of users
func GetUsersAvatarsByIds(c *gin.Context) {

	var usersIdsParam schemas.UsersIdsParam

	err := c.ShouldBindQuery(&usersIdsParam)
	if err != nil {
		log.Println("Bad users ids")
		c.JSON(http.StatusNotFound, gin.H{"detail": "Bad users ids format"})
		return
	}
	usersIds := utils.StrToArrayInt(usersIdsParam.UsersIds)
	avatarsPaths := db.GetUsersAvatarsByIds(usersIds)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"detail": "Image not found"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"avatars": avatarsPaths})
	return
}
