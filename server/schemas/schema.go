package schemas

import "mime/multipart"

type SetAvatarRequestBody struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}
