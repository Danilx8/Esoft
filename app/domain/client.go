package domain

type Client struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	LastName    string `gorm:"column:last_name"`
	FirstName   string `gorm:"column:first_name"`
	MiddleName  string `gorm:"column:middle_name"`
	PhoneNumber string `gorm:"column:phone_number"`
	Email       string `gorm:"column:email"`
}

type ClientRepository interface {
	Create(client *Client) error
	GetClients(clients *[]Client) error
	UpdateClient(client *Client) error
	DeleteClient(client *Client) error
	GetClientById(id int64) error
}
