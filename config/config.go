package config

import (
	"context"
	"fmt"
	"gofirestore/globals"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Connection() (context.Context, *firebase.App) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(globals.GetEnv("JSONKEY"))
	config := &firebase.Config{ProjectID: globals.GetEnv("PROJECTID")}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		fmt.Println("error cuiiiiiiiii")
		log.Fatalln(err)
	}
	return ctx, app
}
