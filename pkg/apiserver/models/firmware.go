package models

import (
	"time"

	"gorm.io/datatypes"
)

type ManufacturerModel struct {
	ID          uint   `gorm:"primaryKey"`
	DisplayName string `gorm:"not null"`
	Status      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ManufacturerModel) TableName() string {
	return "t_manufacturers"
}

func (mm *ManufacturerModel) QueryManufacturerByID(id uint) *ManufacturerModel {
	DB.First(&mm, id)
	return mm
}

type FirmwareModel struct {
	ID           uint `gorm:"primaryKey"`
	Manufacturer uint
	Model        uint
	Version      string
	Search       string
	Status       uint
	FileExtra    datatypes.JSON
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (FirmwareModel) TableName() string {
	return "t_firmwares"
}

func (fm FirmwareModel) QueryFirmwaresByModelID(id uint) []*FirmwareModel {
	fms := []*FirmwareModel{}
	DB.Where("model = ?", id).Find(&fms)
	return fms
}

type ModelModel struct {
	ID           uint   `gorm:"primaryKey"`
	DisplayName  string `gorm:"not null"`
	Manufacturer uint
	Status       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (ModelModel) TableName() string {
	return "t_models"
}

func (mm *ModelModel) QueryModelByID(id uint) *ModelModel {
	DB.First(&mm, id)
	return mm
}
