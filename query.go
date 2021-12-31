package main

import (
	"context"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func findData(collection string) []bson.M {
	data := make([]bson.M, 0)
	if err := MainDB.Collection(collection).Find(context.TODO(), bson.M{}).Limit(3).All(&data); err != nil {
		fmt.Println(err.Error())
	}
	return data

}
