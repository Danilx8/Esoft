package route

import (
	"esoft/app/api/controller"

	"github.com/gin-gonic/gin"
)

func NewPropertyRouter(propertyController *controller.PropertyController, group *gin.RouterGroup) {
	group.POST("properties", propertyController.CreateProperty)
	group.GET("properties", propertyController.GetProperties)
	group.PUT("properties", propertyController.UpdateProperty)
	group.DELETE("properties", propertyController.DeleteProperty)
}
