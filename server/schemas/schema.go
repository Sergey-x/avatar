package schemas

import "mime/multipart"

type SetAvatarRequestBody struct {
	ID     uint64                `form:"id" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}
