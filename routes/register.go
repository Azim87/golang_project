package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"strconv"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user or event!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registered!"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	var event models.Event
	event.Id = eventId

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Cancelled"})
}
