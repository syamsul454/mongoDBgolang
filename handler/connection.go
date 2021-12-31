package handler

import (
	"context"

	"github.com/qiniu/qmgo"
)

var mainDB *qmgo.Database

func init() {
	initMongoDB()
}

func initMongoDB() {
	if client, err := qmgo.NewClient(context.TODO(), &qmgo.Config{Uri: "mongodb://localhost:27017"}); err != nil {
		panic(err)
	} else {
		mainDB = client.Database("latinan")
	}
}
