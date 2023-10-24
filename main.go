package main

import (
	"cloudisk/db"
	"cloudisk/handler"
	"log"
	"net/http"
)

func main() {
	//fmt.Print("hello")
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	//http.Handle("/static/view/home.html", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/file/upload", handler.AuthenInterceptor(handler.UploadHandler))
	http.HandleFunc("/file/upload/suc", handler.AuthenInterceptor(handler.UploadSucHandler))
	//handler.Set()
	//http.HandleFunc("/file/delete", handler.HttpInterceptor())
	log.Print("server start")
	db.InitConnect()
	err := http.ListenAndServe(":8080", nil)
	log.Print(err)

}
