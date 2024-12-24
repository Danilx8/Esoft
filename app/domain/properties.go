package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type PropertyType string

const (
	Apartment PropertyType = "apartment"
	House     PropertyType = "house"
	Land      PropertyType = "land"
)

// TODO: убрать поля для использования в json формате
type Property struct {
	ID            int64           `gorm:"primaryKey;autoIncrement"`
	PropertyType  PropertyType    `gorm:"type:property_type"`
	AddressId     int             `gorm:"column:address_id"`
	CoordinatesId int             `gorm:"column:coordinates_id"`
	Floor         int             `gorm:"column:floor"`
	RoomsCount    int             `gorm:"column:rooms_count"`
	Area          decimal.Decimal `gorm:"column:area"`
	HouseFloors   int             `gorm:"column:house_floors"`
	CreatedAt     *time.Time      `gorm:"column:created_at;type:datetime"`
	UpdatedAt     *time.Time      `gorm:"column:updated_at;type:datetime"`
}

type PropertyRepository interface {
	Create(property *Property) error
	GetProperties(*[]Property) error
	UpdateProperty(property *Property) error
	DeleteProperty(property *Property) error
	GetPropertyById(id int64) error
}
