package handlers

import (
	"avatar.com/avatar/db"
	"avatar.com/avatar/server/conf"
	"avatar.com/avatar/server/schemas"
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
	dirPath := conf.BaseAvatarPath + "/" + dirPrefix
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fullPathToSaveFile := dirPath + filename

	// save on disk
	err = c.SaveUploadedFile(requestBody.Avatar, fullPathToSaveFile)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// save in db
	err = db.SetSrcPath(requestBody.ID, dirPrefix+filename)
	if err != nil {
		log.Println(err.Error())
		return
	}

	responseStatus = http.StatusOK
	responseBody = gin.H{"avatarSrc": fullPathToSaveFile}
}
