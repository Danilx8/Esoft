package domain

type Agent struct {
	ID              int64  `gorm:"primaryKey;autoIncrement"`
	LastName        string `gorm:"column:last_name"`
	FirstName       string `gorm:"column:first_name"`
	MiddleName      string `gorm:"column:middle_name"`
	CommissionShare int64  `gorm:"column:commission_share"`
}

func (Agent) TableName() string {
	return "realtor"
}

type AgentRepository interface {
	Create(client *Agent) error
	GetAgents(clients *[]Agent) error
	UpdateAgent(client *Agent) error
	DeleteAgent(client *Agent) error
	GetAgentById(id int64) error
}
