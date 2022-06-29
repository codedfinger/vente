package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/codedfinger/vente/models"
	"github.com/codedfinger/vente/controllers" 
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	models.ConnectDatabase() //connects to the db function

	r.GET("/events", controllers.GetEvents) // get all event route
	r.POST("/new/event", controllers.AddEvent) // add new event
	r.GET("/event/:id", controllers.GetEvent) //find one event
	r.PATCH("/event/:id", controllers.UpdateEvent) //update event
	r.DELETE("/event/:id", controllers.RemoveEvent) //delete event


	r.Run()
}