package handlers

import (
	"avatar.com/avatar/avatar/config"
	"avatar.com/avatar/avatar/db"
	"avatar.com/avatar/avatar/server/schemas"
	"avatar.com/avatar/avatar/utils"
	"github.com/gin-gonic/gin"
	nanoid "github.com/matoous/go-nanoid/v2"
	"log"
	"net/http"
	"os"
)

// SetAvatar Set user's avatar
func SetAvatar(c *gin.Context) {
	responseStatus := http.StatusBadRequest
	responseBody := gin.H{"error": "couldn't save image"}

	defer func() {
		c.JSON(responseStatus, responseBody)
	}()

	// get data from request multipart form
	var requestBody schemas.SetAvatarRequestBody
	err := c.ShouldBind(&requestBody)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// create filename and full path
	filename, err := nanoid.New()
	if err != nil {
		log.Println(err.Error())
		return
	}

	dirPrefix := filename[0:2] + "/" + filename[2:4] + "/"
	dirPath := config.BaseAvatarPath + "/" + dirPrefix
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fullPathToSaveFile := dirPath + filename

	// get user id from http header
	var userId uint64 = utils.GetUserIdFromHeader(c)
	if userId == 0 {
		log.Println("User id = 0")
		return
	}

	// save on disk
	err = c.SaveUploadedFile(requestBody.Avatar, fullPathToSaveFile)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// save in db
	err = db.SetSrcPath(userId, dirPrefix+filename)
	if err != nil {
		log.Println(err.Error())
		return
	}

	responseStatus = http.StatusOK
	responseBody = gin.H{"avatarSrc": fullPathToSaveFile}
}
