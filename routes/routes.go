package routes

import (
	"github.com/gin-gonic/gin"
)

func InitialiseRoutes(router *gin.Engine) {
	// router.Use(auth.SetUserStatus())
	router.GET("/", ShowHomePage) // Handle the index route
	router.GET("/home", ShowHomePage)
	router.POST("/bot-conversation", BotConversation)
}
