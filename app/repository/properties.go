package repository

import (
	"esoft/app/domain"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type propertyRepository struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) domain.PropertyRepository {
	return &propertyRepository{
		db: db,
	}
}

func (u *propertyRepository) Create(property *domain.Property) error {
	result := u.db.Create(property)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *propertyRepository) GetProperties(propertys *[]domain.Property) error {
	result := u.db.Find(&propertys)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *propertyRepository) UpdateProperty(property *domain.Property) error {
	propertyOld := &domain.Property{}
	result := u.db.Where("ID = ?", property.ID).First(propertyOld)

	if result.Error != nil {
		return fmt.Errorf("failed to fetch property with id %d: %w", property.ID, result.Error)
	}
	propertyVal := reflect.ValueOf(property).Elem()
	propertyOldVal := reflect.ValueOf(propertyOld).Elem()

	for i := 0; i < propertyVal.NumField(); i++ {
		value := propertyVal.Field(i)
		if !value.IsValid() || propertyVal.Type().Field(i).Name == "ID" || value.IsZero() {
			continue
		}
		propertyOldVal.Field(i).Set(value)
	}

	result = u.db.Save(propertyOld)
	if result.Error != nil {
		return fmt.Errorf("failed to update property with id %d: %w", property.ID, result.Error)
	}

	return nil
}

func (u *propertyRepository) GetPropertyById(id int64) error {
	var property domain.Property
	result := u.db.Where("ID = ?", id).First(&property)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *propertyRepository) DeleteProperty(property *domain.Property) error {
	if err := u.GetPropertyById(property.ID); err != nil {
		return err
	}
	result := u.db.Where("ID = ?", property.ID).Delete(&property)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
