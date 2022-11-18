package server

import (
	handlers2 "avatar.com/avatar/avatar/server/handlers"
	teamAvatars2 "avatar.com/avatar/avatar/server/handlers/teamAvatars"
	"avatar.com/avatar/avatar/server/handlers/usersAvatars"
	"fmt"
	"github.com/gin-gonic/gin"
)

func setRoutes(router *gin.Engine) *gin.Engine {
	teamAvatarsPrefix := fmt.Sprintf("avatars/teams/:%s", teamAvatars2.TeamIdParamName)
	usersAvatarsPrefix := "/avatar"

	v1 := router.Group("")
	{
		v1.GET(usersAvatarsPrefix, handlers2.GetAvatarPath)
		v1.POST(usersAvatarsPrefix, handlers2.SetAvatar)
		v1.DELETE(usersAvatarsPrefix, handlers2.DeleteAvatar)

		v1.GET(teamAvatarsPrefix, teamAvatars2.GetTeamAvatarPath)
		v1.POST(teamAvatarsPrefix, teamAvatars2.SetTeamAvatar)
		v1.DELETE(teamAvatarsPrefix, teamAvatars2.DeleteTeamAvatar)

		v1.GET(usersAvatarsPrefix+"/list", usersAvatars.GetUsersAvatarsByIds)
	}

	return router
}
