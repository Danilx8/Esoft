package route

import (
	"esoft/app/api/controller"

	"github.com/gin-gonic/gin"
)

func NewAgentRouter(agentController *controller.AgentController, group *gin.RouterGroup) {
	group.POST("agents", agentController.CreateAgent)
	group.GET("agents", agentController.GetAgents)
	group.PUT("agents", agentController.UpdateAgent)
	group.DELETE("agents", agentController.DeleteAgent)
	group.POST("agents/search", agentController.SearchAgents)
}
