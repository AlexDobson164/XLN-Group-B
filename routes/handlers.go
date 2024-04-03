package routes

import (
	"APR/db"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowHomePage(c *gin.Context) {
	ID, err := db.InsertChatLogEntry()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusForbidden)
	}
	c.SetCookie("ID", strconv.Itoa(ID), 3600, "", "", false, true)
	Render(c, gin.H{
		"title": "Home",
	}, "Home.html")
}

type T struct {
	Message string `json:"Message"`
	IsBot   bool   `json:"IsBot"`
}

func BotConversation(c *gin.Context) {
	ID, e := c.Cookie("ID")
	if e != nil {
		log.Print(e)
		return
	}
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
		return
	}
	a := new(T)
	json.Unmarshal(jsonData, a)
	log.Print(string(jsonData))
	log.Print(a)
	log.Print(a.Message)
	NID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	db.InsertBotConversation(NID, a.Message, a.IsBot)
	c.Status(200)
}
