package itemHandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lorduwahz/pular/database"
	"github.com/lorduwahz/pular/schema"
)

var db = database.DbConfig()

func GetItems(c *gin.Context) {
	var items []schema.Item

	db.Find(&items)
	c.IndentedJSON(http.StatusOK, items)

}

func PostItem(c *gin.Context) {
	var item schema.Item

	err := c.BindJSON(&item)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	db.Create(&item)
	c.JSON(http.StatusOK, item)
}
