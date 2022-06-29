package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/codedfinger/vente/models"
)

type CreateEventInput struct {
	Title	string	`json:"title" binding:"required"`
	Venue	string	`json:"venue" binding:"required"`
	Price	string	`json:"price" binding:"required"`
}

type UpdateEventInput struct {
	Title	string	`json:"title"`
	Venue	string	`json:"venue"`
	Price	string	`json:"price"`
}

// GET /events
// get all events
func GetEvents(c *gin.Context) {
	var events []models.Event
	models.DB.Find(&events)

	c.JSON(http.StatusOK, gin.H{"data": events})
}

// POST /new/event
// create new event
func AddEvent(c *gin.Context){
	//validate inpput
	var input CreateEventInput
	if err := c.ShouldBindJSON(&input); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create event
	event := models.Event{Title: input.Title, Venue: input.Venue, Price: input.Price}
	models.DB.Create(&event)

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// GET /event/:id
// Find an event
func GetEvent (c *gin.Context) {
	var event models.Event

	if err := models.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"eror": "Record not found"})
	}

	c.JSON(http.StatusOK, gin.H{"data": event})


}

// PATCH /event/:id
// update book
func UpdateEvent(c *gin.Context) {
	// Get model if exist
	var event models.Event
	if err := models.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	//validate input
	var input UpdateEventInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&event).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// DELETE /event/:id
// Removve an event
func RemoveEvent (c *gin.Context) {
	// Get model if exist
	var event models.Event
	if err := models.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&event)

	c.JSON(http.StatusOK, gin.H{"data": "item deleted successfully"})
}