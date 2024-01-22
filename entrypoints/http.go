package entrypoints

import (
	"github.com/Evefest/infrastructureEventWrite/handlers"
	"github.com/gin-gonic/gin"
)

func SetEntryPoints(api *gin.RouterGroup) {
	api.GET("/", handlers.WelcomeApi)
}
