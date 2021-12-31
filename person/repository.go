package person

import (
	"context"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

// struct person
type Repository interface {
	FindAll([]bson.M, error)
	GroupByTypeOfName([]bson.M, error)
}

type repository struct {
	db *qmgo.Database
}

func NewRepository(db *qmgo.Database) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]bson.M, error) {
	var person []bson.M
	err := r.db.Collection("person").Find(context.TODO(), bson.M{}).All(&person)
	return person, err
}

func (r *repository) GroupByTypeOfName() ([]bson.M, error) {
	filter := []bson.M{{"$group": bson.M{"_id": "$Type of name", "count": bson.M{"$sum": 1}, "genre": bson.M{"$first": "$Name"}}}}
	var person []bson.M
	err := r.db.Collection("person").Aggregate(context.Background(), filter).All(&person)
	return person, err
}
