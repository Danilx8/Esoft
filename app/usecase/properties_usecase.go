package usecase

import (
	"esoft/app/domain"

	"github.com/gin-gonic/gin"
)

type PropertyUsecase struct {
	PropertyRepository domain.PropertyRepository
}

func NewPropertyUsecase(clientRepo domain.PropertyRepository) *PropertyUsecase {
	return &PropertyUsecase{
		PropertyRepository: clientRepo,
	}
}

func (u *PropertyUsecase) CreateProperty(c *gin.Context, client *domain.Property) error {
	err := u.PropertyRepository.Create(client)
	if err != nil {
		return err
	}
	return err
}

func (u *PropertyUsecase) GetProperties(c *gin.Context) (*[]domain.Property, error) {
	var clients []domain.Property
	err := u.PropertyRepository.GetProperties(&clients)
	if err != nil {
		return nil, err
	}
	return &clients, nil
}

func (u *PropertyUsecase) UpdateProperty(c *gin.Context, client *domain.Property) error {
	err := u.PropertyRepository.UpdateProperty(client)
	if err != nil {
		return err
	}
	return nil
}

func (u *PropertyUsecase) DeleteProperty(c *gin.Context, client *domain.Property) error {
	err := u.PropertyRepository.DeleteProperty(client)
	if err != nil {
		return err
	}
	return nil
}

//func (u *PropertyUsecase) SearchSimilaryProperties(c *gin.Context, property *domain.Property) (*[]domain.Property, error) {
//	var propertys []domain.Property
//	err := u.PropertyRepository.GetProperties(&propertys)
//	if err != nil {
//		return nil, err
//	}
//	target := internal.ComposerFLM(property., property.LastName, property.MiddleName)
//	foundProperties := make([]domain.Property, 0, len(propertys)/8)
//	for i := range propertys {
//		tmpInput := internal.ComposerFLM(propertys[i].FirstName, propertys[i].LastName, propertys[i].MiddleName)
//		distance := levenshtein.ComputeDistance(target, tmpInput)
//		if distance <= 3 {
//			foundProperties = append(foundProperties, propertys[i])
//		}
//	}
//
//	return &foundProperties, nil
//}
