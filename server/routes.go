package server

import (
	"avatar.com/avatar/server/handlers"
	"avatar.com/avatar/server/handlers/teamAvatars"
	"avatar.com/avatar/server/handlers/usersAvatars"
	"fmt"
	"github.com/gin-gonic/gin"
)

func setRoutes(router *gin.Engine) *gin.Engine {
	teamAvatarsPrefix := fmt.Sprintf("avatar/teams/:%s", teamAvatars.TeamIdParamName)
	usersAvatarsPrefix := "/avatar"

	v1 := router.Group("")
	{
		v1.GET(usersAvatarsPrefix, handlers.GetAvatarPath)
		v1.POST(usersAvatarsPrefix, handlers.SetAvatar)
		v1.DELETE(usersAvatarsPrefix, handlers.DeleteAvatar)

		v1.GET(teamAvatarsPrefix, teamAvatars.GetTeamAvatarPath)
		v1.POST(teamAvatarsPrefix, teamAvatars.SetTeamAvatar)
		v1.DELETE(teamAvatarsPrefix, teamAvatars.DeleteTeamAvatar)

		v1.GET(usersAvatarsPrefix+"/list", usersAvatars.GetUsersAvatarsByIds)
		v1.GET("/avatar/list/team", teamAvatars.GetTeamsAvatarsByIds)
	}

	return router
}
