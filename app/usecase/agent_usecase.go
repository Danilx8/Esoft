package usecase

import (
	"esoft/app/domain"
	"esoft/app/internal"

	"github.com/agnivade/levenshtein"
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

func (u *AgentUsecase) SearchSimilaryAgents(c *gin.Context, agent *domain.Agent) (*[]domain.Agent, error) {
	var agents []domain.Agent
	err := u.AgentRepository.GetAgents(&agents)
	if err != nil {
		return nil, err
	}
	target := internal.ComposerFLM(agent.FirstName, agent.LastName, agent.MiddleName)
	foundAgents := make([]domain.Agent, 0, len(agents)/8)
	for i := range agents {
		tmpInput := internal.ComposerFLM(agents[i].FirstName, agents[i].LastName, agents[i].MiddleName)
		distance := levenshtein.ComputeDistance(target, tmpInput)
		if distance <= 3 {
			foundAgents = append(foundAgents, agents[i])
		}
	}

	return &foundAgents, nil
}
