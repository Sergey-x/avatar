package db

func GetTeamSrcPath(id uint64) (avatarSrcPath string, err error) {
	resultRow := db.QueryRow("SELECT (src_path) FROM "+TeamAvatarTableName+" WHERE team_id = $1;", id)
	err = resultRow.Scan(&avatarSrcPath)
	return
}

func SetTeamSrcPath(id uint64, srcPath string) (err error) {
	_, err = db.Exec("INSERT INTO "+TeamAvatarTableName+"(team_id, src_path) VALUES"+
		"($1, $2) ON CONFLICT (team_id) DO UPDATE SET src_path = $2;", id, srcPath)
	return
}

func DeleteTeamSrcPath(id uint64) (err error) {
	result := db.QueryRow("DELETE FROM "+TeamAvatarTableName+" WHERE team_id = $1;", id)
	err = result.Err()
	return err
}
