package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type TeamAvatar struct {
	TeamId  int    `json:"teamId"`
	SrcPath string `json:"srcPath"`
}

func GetTeamsAvatarsByIds(ids []int) (teamsAvatars []TeamAvatar) {
	teamsIds := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")
	rows, err := db.Query(fmt.Sprintf("SELECT team_id, src_path FROM %s WHERE team_id IN (%s);", TeamAvatarTableName, teamsIds))
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
		teamAvatar := TeamAvatar{}
		err := rows.Scan(&teamAvatar.TeamId, &teamAvatar.SrcPath)
		if err != nil {
			return nil
		}
		teamsAvatars = append(teamsAvatars, teamAvatar)
	}
	log.Println(teamsAvatars)
	j, _ := json.Marshal(&teamsAvatars)
	log.Println(string(j))
	return teamsAvatars
}
