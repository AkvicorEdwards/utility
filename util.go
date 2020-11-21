package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"util/def"
	"util/handler"
	"util/tpl"
)

func main() {
	tpl.Parse()
	handler.ParsePrefix()
	addr := fmt.Sprintf(":%d", def.PORT)
	server := http.Server{
		Addr:              addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Minute,
	}
	log.Printf("http://127.0.0.1%s", addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}