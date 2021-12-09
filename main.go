package main

import (
	"log"

	"github.com/gin-gonic/gin"
	itemHandler "github.com/lorduwahz/pular/handlers/item-handler"
	userHandlers "github.com/lorduwahz/pular/handlers/user-handler"
	"github.com/lorduwahz/pular/schema"
)

func main() {
	schema.Migrations()

	router := gin.Default()
	userGroup := router.Group("user")
	{
		userGroup.POST("/register", userHandlers.PostUser)
		userGroup.GET("/", userHandlers.GetUsers)
		userGroup.GET("/:id", userHandlers.GetUser)
		userGroup.DELETE("/delete/:id", userHandlers.DeleteUser)
		userGroup.PUT("/update/:id", userHandlers.UpdateUser)
		userGroup.GET("/getitems", userHandlers.GetUserandItems)
	}

	itemGroup := router.Group("item")
	{
		itemGroup.POST("/add-item", itemHandler.PostItem)
		itemGroup.GET("/", itemHandler.GetItems)
	}

	log.Fatal(router.Run("localhost:8080"))
	// router.Run("localhost:8080")
}
