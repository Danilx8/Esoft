package route

import (
	"esoft/app/api/controller"

	"github.com/gin-gonic/gin"
)

func NewClientRouter(clientController *controller.ClientController, group *gin.RouterGroup) {
	group.POST("clients", clientController.CreateClient)
	group.GET("clients", clientController.GetClients)
	group.PUT("clients", clientController.UpdateClient)
	group.DELETE("clients", clientController.DeleteClient)
}
