package user

import "time"

type User struct {
	Name      string    `bson:"name"`
	UserName  string    `bson:"user_name"`
	Password  string    `bson:"password"`
	Email     string    `bson:"email"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

var StringData string = "hello"
