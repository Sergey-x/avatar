package teamAvatars

import (
	"avatar.com/avatar/db"
	"avatar.com/avatar/server/conf"
	"avatar.com/avatar/server/schemas"
	"avatar.com/avatar/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	nanoid "github.com/matoous/go-nanoid/v2"
	"log"
	"net/http"
	"os"
	"strconv"
)

// SetTeamAvatar Set team's avatar
func SetTeamAvatar(c *gin.Context) {
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

	// get user id from http header
	var userId uint64 = utils.GetUserIdFromHeader(c)
	if userId == 0 {
		log.Println("User id = 0")
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
		responseStatus = http.StatusForbidden
		responseBody = gin.H{"Detail": "Permission denied"}
		c.JSON(http.StatusForbidden, gin.H{"Detail": "Permission denied"})
		return
	}

	// save on disk
	err = c.SaveUploadedFile(requestBody.Avatar, fullPathToSaveFile)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// save in db
	err = db.SetTeamSrcPath(teamId, dirPrefix+filename)
	if err != nil {
		log.Println(err.Error())
		return
	}

	responseStatus = http.StatusOK
	responseBody = gin.H{"avatarSrc": fullPathToSaveFile}
}
