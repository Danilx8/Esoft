package repository

import (
	"esoft/app/domain"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type agentRepository struct {
	db *gorm.DB
}

func NewAgentRepository(db *gorm.DB) domain.AgentRepository {
	return &agentRepository{
		db: db,
	}
}

func (u *agentRepository) Create(agent *domain.Agent) error {
	result := u.db.Create(agent)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *agentRepository) GetAgents(agents *[]domain.Agent) error {
	result := u.db.Find(&agents)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *agentRepository) UpdateAgent(agent *domain.Agent) error {
	agentOld := &domain.Agent{}
	result := u.db.Where("ID = ?", agent.ID).First(agentOld)

	if result.Error != nil {
		return fmt.Errorf("failed to fetch agent with id %d: %w", agent.ID, result.Error)
	}
	agentVal := reflect.ValueOf(agent).Elem()
	agentOldVal := reflect.ValueOf(agentOld).Elem()

	for i := 0; i < agentVal.NumField(); i++ {
		value := agentVal.Field(i)
		if !value.IsValid() || agentVal.Type().Field(i).Name == "ID" {
			continue
		}
		agentOldVal.Field(i).Set(value)
	}

	result = u.db.Save(agentOld)
	if result.Error != nil {
		return fmt.Errorf("failed to update agent with id %d: %w", agent.ID, result.Error)
	}

	return nil
}

func (u *agentRepository) GetAgentById(id int64) error {
	var agent domain.Agent
	result := u.db.Where("ID = ?", id).First(&agent)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *agentRepository) DeleteAgent(agent *domain.Agent) error {
	if err := u.GetAgentById(agent.ID); err != nil {
		return err
	}
	result := u.db.Where("ID = ?", agent.ID).Delete(&agent)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
