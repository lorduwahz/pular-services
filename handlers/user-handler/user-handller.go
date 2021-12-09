package userHandlers

import (
	"fmt"
	"strconv"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lorduwahz/pular/database"
	"github.com/lorduwahz/pular/schema"
)

var db = database.DbConfig()

// Get all users in the database
func GetUsers(c *gin.Context) {
	var users []*schema.User

	db.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}

// Get a single user form their ID
func GetUser(c *gin.Context) {
	var user schema.User
	id := c.Param("id")
	db.First(&user, id)
	c.IndentedJSON(http.StatusOK, user)
}

type result struct {
	full_name string
	item_name string
}

func GetUserandItems(c *gin.Context) {
	var user schema.User
	var result result

	db.Model(&user).Select("users.full_name, items.name").Joins("left join items on items.user_id = users.id").Scan(&result)
	// c.BindJSON(&result)
	fmt.Println(result)

	//user.Items = items
	c.IndentedJSON(http.StatusOK, result)

}

// Post a new user to the databse
func PostUser(c *gin.Context) {
	var user schema.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	db.Create(&user)
	c.JSON(http.StatusOK, user)

}

// Update an already existing user
func UpdateUser(c *gin.Context) {
	var singleUser schema.User
	id, _ := strconv.Atoi(c.Param(("id")))

	// check if db.First returns an err eg. result not found.
	// if it does, err catches it and the program exits with err message
	// if there is nor err, it goes ahead to update the user
	err := db.First(&singleUser, id).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user with such id"})
	} else {
		c.BindJSON(&singleUser)
		db.Save(&singleUser)
		c.IndentedJSON(http.StatusOK, singleUser)
	}

}

// Permanently delete a user from the database.
func DeleteUser(c *gin.Context) {
	var user schema.User
	id := c.Param("id")

	//db.Delete(&user, id) => soft delete
	//keeps info in the database and add the time the delete command as called.

	//db.Unscoped().Delete deletes the info from the database permanently.
	db.Unscoped().Delete(&user, id)
	c.IndentedJSON(http.StatusOK, user)

}
