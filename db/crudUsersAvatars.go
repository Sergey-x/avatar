package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type UserAvatar struct {
	UserId  int    `json:"userId"`
	SrcPath string `json:"srcPath"`
}

func GetUsersAvatarsByIds(ids []int) (usersAvatars []UserAvatar) {
	usersIds := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")
	var unknownIds = map[int]bool{}
	for userId, _ := range usersIds {
		unknownIds[userId] = false
	}

	rows, err := db.Query(fmt.Sprintf("SELECT user_id, src_path FROM %s WHERE user_id IN (%s);", UserAvatarTableName, usersIds))

	if err != nil {
		err := rows.Close()
		if err != nil {
			return nil
		}
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		userAvatar := UserAvatar{}
		err := rows.Scan(&userAvatar.UserId, &userAvatar.SrcPath)
		if err != nil {
			return nil
		}
		usersAvatars = append(usersAvatars, userAvatar)
		unknownIds[userAvatar.UserId] = true
	}

	for userId, _ := range usersIds {
		if unknownIds[userId] == false {
			userAvatar := UserAvatar{}
			userAvatar.UserId = userId
			userAvatar.SrcPath = ""
			if err != nil {
				return nil
			}
			usersAvatars = append(usersAvatars, userAvatar)
		}
	}

	log.Println(usersAvatars)
	j, _ := json.Marshal(&usersAvatars)
	log.Println(string(j))
	return usersAvatars
}
