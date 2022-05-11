package db

func GetSrcPath(userId uint64) (avatarSrcPath string, e error) {
	resultRow := db.QueryRow("SELECT (src_path) FROM avatar WHERE user_id = $1;", userId)
	e = resultRow.Scan(&avatarSrcPath)
	return

}

func SetSrcPath(userId uint64, srcPath string) (err error) {
	_, err = db.Exec("INSERT INTO avatar (user_id, src_path) VALUES"+
		"($1, $2) ON CONFLICT (user_id) DO UPDATE SET src_path = $2;", userId, srcPath)
	return
}
