package handlers

import (
	"github.com/Evefest/domainEventWrite/constants"
	"github.com/Evefest/domainEventWrite/exceptions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HomeResponse struct {
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

func WelcomeApi(context *gin.Context) {
	context.JSON(200, HomeResponse{Time: time.Now(), Message: constants.WelcomeMessage})
}

func validateRequest(context *gin.Context, request any) (bool, exceptions.Response) {
	if err := context.ShouldBindJSON(&request); err != nil {
		return false, exceptions.Response{}.Create(http.StatusBadRequest, constants.BadRequest)
	}
	return true, exceptions.Response{}
}
