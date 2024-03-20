package routes

import (
	"github.com/gin-gonic/gin"
)

func InitialiseRoutes(router *gin.Engine) {
	// router.Use(auth.SetUserStatus())
	router.GET("/", ShowIndexPage) // Handle the index route
	router.GET("/nutrition", ShowNutritionPage)
	router.GET("/pin", ShowPinPage)
	router.GET("/coaching", ShowCoachingPage)
	router.GET("/home", ShowHomePage)
	// logRoutes := router.Group("/personal-log") // log group
	// {
	// 	logRoutes.GET("/view", auth.EnsureLoggedIn(), ShowLogs)
	// 	logRoutes.GET("/view/:log_id", auth.EnsureLoggedIn(), GetLog)
	// 	logRoutes.GET("/create", auth.EnsureLoggedIn(), ShowLogCreationPage)
	// 	logRoutes.POST("/create", auth.EnsureLoggedIn(), CreateLog)
	// }
}