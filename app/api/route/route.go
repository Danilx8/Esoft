package route

import (
	"esoft/app/api/controller"
	"esoft/app/repository"
	"esoft/app/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, e *gin.Engine) {
	clientRepository := repository.NewClientRepository(db)

	clientController := controller.ClientController{
		ClientUsecase: usecase.NewClientUsecase(clientRepository),
	}

	publicRouter := e.Group("")
	NewClientRouter(&clientController, publicRouter)
	SwaggerRouter(publicRouter)
}
