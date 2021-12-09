package schema

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lorduwahz/pular/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Items    []Item `gorm:"foreignKey:UserID"`
}

type Item struct {
	gorm.Model
	Name        string `json:"name"`
	Decsription string `json:"description"`
	UserID      uint
}

func Migrations() {
	db := database.DbConfig()

	db.AutoMigrate(&User{}, &Item{})
	fmt.Println("Successfully made DB migrations.")
}

func PostUser(c *gin.Context) {
	db := database.DbConfig()

	var user User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	db.Create(&user)
	c.IndentedJSON(http.StatusOK, user)
}
