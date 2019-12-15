package main

import (
	"gofirestore/globals"
	"gofirestore/routes"
	"net/http"
	"time"
)

func main() {
	r := routes.SetupRouter()
	s := &http.Server{
		Addr:           globals.GetEnv("PORT"),
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
