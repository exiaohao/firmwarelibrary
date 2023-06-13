package models

import (
	"encoding/json"
	"log"
	"time"
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

type FileExtraStruct struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
	Size uint64 `json:"size"`
}

type FirmwareModel struct {
	ID                      uint `gorm:"primaryKey"`
	Manufacturer            uint
	ManufacturerDisplayName string `gorm:"-"`
	Model                   uint
	ModelDisplayName        string `gorm:"-"`
	Version                 string
	Search                  string
	Status                  uint
	FileExtra               []byte          `json:"-"`
	FileExtraJson           FileExtraStruct `gorm:"-",json:"file_extra"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func (FirmwareModel) TableName() string {
	return "t_firmwares"
}

func (fm FirmwareModel) QueryFirmwaresByModelID(id uint) []*FirmwareModel {
	fms := []*FirmwareModel{}
	DB.Where("model = ?", id).Find(&fms)
	return fms
}

func (fm *FirmwareModel) QueryFirmwareByID(id uint) {
	DB.First(&fm, id)
}

func (fm *FirmwareModel) FillRelated() error {
	mm := ModelModel{}
	mm.QueryModelByID(fm.Model)
	mam := ManufacturerModel{}
	mam.QueryManufacturerByID(fm.Manufacturer)
	fm.ModelDisplayName = mm.DisplayName
	fm.ManufacturerDisplayName = mam.DisplayName
	fe := new(FileExtraStruct)
	if err := json.Unmarshal(fm.FileExtra, fe); err != nil {
		log.Println(fm.ID, fm.FileExtra)
		log.Fatalln(err)
		return err
	}
	fm.FileExtraJson = *fe
	return nil
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
