package usecase

import (
	"esoft/app/domain"

	"github.com/gin-gonic/gin"
)

type AgentUsecase struct {
	AgentRepository domain.AgentRepository
}

func NewAgentUsecase(clientRepo domain.AgentRepository) *AgentUsecase {
	return &AgentUsecase{
		AgentRepository: clientRepo,
	}
}

func (u *AgentUsecase) CreateAgent(c *gin.Context, client *domain.Agent) error {
	err := u.AgentRepository.Create(client)
	if err != nil {
		return err
	}
	return err
}

func (u *AgentUsecase) GetAgents(c *gin.Context) (*[]domain.Agent, error) {
	var clients []domain.Agent
	err := u.AgentRepository.GetAgents(&clients)
	if err != nil {
		return nil, err
	}
	return &clients, nil
}

func (u *AgentUsecase) UpdateAgent(c *gin.Context, client *domain.Agent) error {
	err := u.AgentRepository.UpdateAgent(client)
	if err != nil {
		return err
	}
	return nil
}

func (u *AgentUsecase) DeleteAgent(c *gin.Context, client *domain.Agent) error {
	err := u.AgentRepository.DeleteAgent(client)
	if err != nil {
		return err
	}
	return nil
}
