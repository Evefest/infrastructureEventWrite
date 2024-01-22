package handlers

import (
	"github.com/Evefest/infrastructureEventWrite/mappers"
	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var eventRequest mappers.EventRequest
	isValid, response := validateRequest(context, &eventRequest)
	if !isValid {
		context.AbortWithStatusJSON(response.ErrorCode, response)
		return
	}
}

func UpdateEvent(context *gin.Context) {

}

func DisableEvent(context *gin.Context) {

}
