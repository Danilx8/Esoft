package usecase

import (
	"esoft/app/domain"
	"esoft/app/internal"

	"github.com/agnivade/levenshtein"
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

func (u *ClientUsecase) SearchSimilaryClients(c *gin.Context, client *domain.Client) (*[]domain.Client, error) {
	var clients []domain.Client
	err := u.ClientRepository.GetClients(&clients)
	if err != nil {
		return nil, err
	}
	target := internal.ComposerFLM(client.FirstName, client.LastName, client.MiddleName)
	foundClients := make([]domain.Client, 0, len(clients)/8)
	for i := range clients {
		tmpInput := internal.ComposerFLM(clients[i].FirstName, clients[i].LastName, clients[i].MiddleName)
		distance := levenshtein.ComputeDistance(target, tmpInput)
		if distance <= 3 {
			foundClients = append(foundClients, clients[i])
		}
	}

	return &foundClients, nil
}
