package usecase

import (
	"esoft/app/domain"

	"github.com/gin-gonic/gin"
)

type ClientUsecase struct {
	ClientRepository domain.ClientRepository
}

func NewClientUsecase(clientRepo domain.ClientRepository) *ClientUsecase {
	return &ClientUsecase{
		ClientRepository: clientRepo,
	}
}

func (u *ClientUsecase) CreateClient(c *gin.Context, client *domain.Client) error {
	err := u.ClientRepository.Create(client)
	if err != nil {
		return err
	}
	return err
}

func (u *ClientUsecase) GetClients(c *gin.Context) (*[]domain.Client, error) {
	var clients []domain.Client
	err := u.ClientRepository.GetClients(&clients)
	if err != nil {
		return nil, err
	}
	return &clients, nil
}

func (u *ClientUsecase) UpdateClient(c *gin.Context, client *domain.Client) error {
	err := u.ClientRepository.UpdateClient(client)
	if err != nil {
		return err
	}
	return nil
}

func (u *ClientUsecase) DeleteClient(c *gin.Context, client *domain.Client) error {
	err := u.ClientRepository.DeleteClient(client)
	if err != nil {
		return err
	}
	return nil
}
