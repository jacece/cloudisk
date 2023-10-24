package handler

import (
	"cloudisk/model"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, header, err := r.FormFile("file")
		if err != nil {
			log.Print("FormFile failed.")
			return
		}
		defer file.Close()
		if err != nil {
			log.Print("r.FormFile error")
			return
		}
		defer file.Close()
		file_meta := model.FileMeta{
			Name:     header.Filename,
			Location: "/usr/share/storage/" + header.Filename,
			Time:     time.Now().Format("2005-01-01 12:11:33"),
		}
		local_file, err := os.Create(file_meta.Location)
		if err != nil {
			log.Print("Create file failed.")
		}
		defer local_file.Close()
		file_meta.Size, err = io.Copy(local_file, file)
		if err != nil {
			log.Print("Failed to save file in local storage.")
			return
		}
		user_file := model.UserFile{
			FileName:     file_meta.Name,
			FileLocation: file_meta.Location,
			CreateTime:   file_meta.Time,
		}

		succ := model.UserFileInsert(user_file)
		if succ {
			http.Redirect(w, r, "/static/view/home.html", http.StatusFound)
		} else {
			w.Write([]byte("Upload Failed!"))
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished")
}
