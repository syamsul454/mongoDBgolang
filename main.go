package main

import (
	"fmt"
	"mongodbGolang/handler"
	"mongodbGolang/user"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(DataString)
	fmt.Println(user.StringData)
	router := gin.Default()
	router.GET("/person", handler.GetPerson)
	router.GET("/person/group", handler.GetPersonGroup)
	router.Run(":8888")

}
