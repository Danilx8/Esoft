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

	clientController := controller.ClientController{
		ClientUsecase: usecase.NewClientUsecase(clientRepository),
	}
	agentController := controller.AgentController{
		AgentUsecase: usecase.NewAgentUsecase(agentRepository),
	}

	publicRouter := e.Group("")
	NewClientRouter(&clientController, publicRouter)
	NewAgentRouter(&agentController, publicRouter)
	SwaggerRouter(publicRouter)
}
