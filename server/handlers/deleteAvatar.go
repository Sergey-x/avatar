package handlers

import (
	"avatar.com/avatar/db"
	"avatar.com/avatar/server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// DeleteAvatar Delete user's avatar
func DeleteAvatar(c *gin.Context) {
	// get user id from http header
	var userId uint64 = utils.GetUserIdFromHeader(c)
	if userId == 0 {
		log.Println("User id = 0")
		c.JSON(http.StatusBadRequest, gin.H{"Detail": "Ошибка пользователя - нераспознанный id"})
		return
	}

	e := db.DeleteSrcPath(userId)
	if e != nil {
		c.JSON(http.StatusNotFound, gin.H{"Detail": "Аватар не сущетсвует"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"Detail": "Аватар удален"})
}
