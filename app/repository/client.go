package repository

import (
	"esoft/app/domain"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) domain.ClientRepository {
	return &clientRepository{
		db: db,
	}
}

func (u *clientRepository) Create(client *domain.Client) error {
	result := u.db.Table("client").Create(client)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *clientRepository) GetClients(clients *[]domain.Client) error {
	result := u.db.Table("client").Find(&clients)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *clientRepository) UpdateClient(client *domain.Client) error {
	clientOld := &domain.Client{}
	result := u.db.Table("client").Where("ID = ?", client.ID).First(clientOld)

	if result.Error != nil {
		return fmt.Errorf("failed to fetch client with id %d: %w", client.ID, result.Error)
	}
	clientVal := reflect.ValueOf(client).Elem()
	clientOldVal := reflect.ValueOf(clientOld).Elem()

	for i := 0; i < clientVal.NumField(); i++ {
		value := clientVal.Field(i)
		if !value.IsValid() || clientVal.Type().Field(i).Name == "ID" {
			continue
		}
		clientOldVal.Field(i).Set(value)
	}

	result = u.db.Table("client").Save(clientOld)
	if result.Error != nil {
		return fmt.Errorf("failed to update client with id %d: %w", client.ID, result.Error)
	}

	return nil
}

func (u *clientRepository) GetClientById(id int64) error {
	var client domain.Client
	result := u.db.Table("client").Where("ID = ?", id).First(&client)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *clientRepository) DeleteClient(client *domain.Client) error {
	if err := u.GetClientById(client.ID); err != nil {
		return err
	}
	result := u.db.Table("client").Where("ID = ?", client.ID).Delete(&client)
	fmt.Println("YRES", client.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
