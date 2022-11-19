package schemas

import (
	"mime/multipart"
)

type SetAvatarRequestBody struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

type UsersIdsParam struct {
	UsersIds string `form:"usersIds" binding:"required"`
}

type TeamsIdsParam struct {
	TeamsIds string `form:"teamsIds" binding:"required"`
}
