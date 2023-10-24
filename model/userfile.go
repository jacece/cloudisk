package model

import (
	"cloudisk/db"
	"log"
)

type UserFile struct {
	FileId       string
	FileName     string
	UserName     string
	FileLocation string
	CreateTime   string
}

func UserFileInsert(userFile UserFile) bool {
	stmt, err := db.DBConn().Prepare(
		"insert ignore into tbl_user_file (`file_name`,`user_name`,`file_loacton`," +
			"`create_time`) values (?,?,?,?,?)")
	if err != nil {
		log.Print("Insert userfile failed.")
		return false
	}

	_, err = stmt.Exec(userFile.FileName, userFile.UserName, userFile.FileLocation, userFile.CreateTime)
	return err == nil
}
