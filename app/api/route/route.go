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
	agentRepository := repository.NewAgentRepository(db)
	propertyRepository := repository.NewPropertyRepository(db)

	clientController := controller.ClientController{
		ClientUsecase: usecase.NewClientUsecase(clientRepository),
	}
	agentController := controller.AgentController{
		AgentUsecase: usecase.NewAgentUsecase(agentRepository),
	}
	propertyController := controller.PropertyController{
		PropertyUsecase: usecase.NewPropertyUsecase(propertyRepository),
	}

	publicRouter := e.Group("")
	NewClientRouter(&clientController, publicRouter)
	NewAgentRouter(&agentController, publicRouter)
	NewPropertyRouter(&propertyController, publicRouter)
	SwaggerRouter(publicRouter)
}
