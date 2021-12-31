package handler

import (
	"fmt"
	"mongodbGolang/person"

	"github.com/gin-gonic/gin"
)

func GetPersonGroup(c *gin.Context) {
	person := person.NewRepository(mainDB)
	data, err := person.GroupByTypeOfName()

	if err != nil {
		fmt.Println("==>", err.Error())
	}

	respon := make([]interface{}, 0)

	for _, v := range data {
		res := map[string]interface{}{
			"name_organisasi": v["_id"],
			"jumlah":          v["count"],
			"genre":           v["genre"],
		}
		respon = append(respon, res)
	}
	c.JSON(200, gin.H{"success": true, "data": respon})

}

func GetPerson(c *gin.Context) {
	person := person.NewRepository(mainDB)
	data, err := person.FindAll()
	if err != nil {
		fmt.Println("==>", err.Error())
	}
	respon := make([]interface{}, 0)
	for _, v := range data {
		res := map[string]interface{}{
			"name":  v["Name"],
			"title": v["Title"],
		}
		respon = append(respon, res)
	}

	c.JSON(200, gin.H{"success": true, "data": respon})

}
