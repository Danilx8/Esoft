package controller

import (
	"esoft/app/domain"
	"esoft/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgentController struct {
	AgentUsecase *usecase.AgentUsecase
}

// CreateAgent godoc
// @Summary	Create of agent
// @Tags Agent
// @Accept json
// @Produce json
// @Param        data    body     domain.Agent true  "scheme of agent"
// @Success 200 {object} domain.Agent
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /agents [post]
func (u AgentController) CreateAgent(c *gin.Context) {
	var agent domain.Agent

	err := c.BindJSON(&agent)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "agent schema", Body: err.Error()})
		return
	}
	if agent.FirstName == "" || agent.LastName == "" || agent.MiddleName == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "agent schema", Body: "Set first name, last name and middle name!"})
		return
	}

	err = u.AgentUsecase.CreateAgent(c, &agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Create new agent",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, agent)
}

// GetAgents godoc
// @Summary	Get of agents
// @Tags Agent
// @Accept json
// @Produce json
// @Success 200 {object} domain.Agent
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /agents [get]
func (u AgentController) GetAgents(c *gin.Context) {
	agents, err := u.AgentUsecase.GetAgents(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Get all agents",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, agents)
}

// UpdateAgent godoc
// @Summary	Update of agent
// @Tags Agent
// @Accept json
// @Produce json
// @Param        data    body     domain.Agent true  "scheme of agent"
// @Success 200 {object} domain.Agent
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /agents [put]
func (u AgentController) UpdateAgent(c *gin.Context) {
	var agent domain.Agent

	err := c.BindJSON(&agent)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "agent schema", Body: err.Error()})
		return
	}

	if agent.FirstName == "" || agent.LastName == "" || agent.MiddleName == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "agent schema", Body: "Set first name, last name and middle name!"})
		return
	}
	err = u.AgentUsecase.UpdateAgent(c, &agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Update of agent",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, agent)
}

// DeleteAgent godoc
// @Summary	Delete of agent
// @Tags Agent
// @Accept json
// @Produce json
// @Param        data    body     domain.Agent true  "scheme of agent"
// @Success 200 {object} domain.Agent
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /agents [delete]
func (u AgentController) DeleteAgent(c *gin.Context) {
	var agent domain.Agent

	err := c.BindJSON(&agent)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "agent schema", Body: err.Error()})
		return
	}

	err = u.AgentUsecase.DeleteAgent(c, &agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Delete of agent",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, agent)
}

// SearchAgent godoc
// @Summary	Search of agent
// @Tags Agent
// @Accept json
// @Produce json
// @Param        data    body     domain.Agent true  "scheme of agent"
// @Success 200 {object} domain.Agent
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /agents/search [post]
func (u AgentController) SearchAgents(c *gin.Context) {
	var agent domain.Agent

	err := c.BindJSON(&agent)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "agent schema", Body: err.Error()})
		return
	}

	agents, err := u.AgentUsecase.SearchSimilaryAgents(c, &agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Search of agent",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, agents)
}
