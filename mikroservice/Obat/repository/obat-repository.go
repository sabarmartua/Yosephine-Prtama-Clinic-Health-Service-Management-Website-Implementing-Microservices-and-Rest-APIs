package repository

import (
	"github.com/sabarmartua/Obat/model"
	"gorm.io/gorm"
)

type ObatRepository interface {
	InsertObat(obat model.Obat) model.Obat
	UpdateObat(obat model.Obat) model.Obat
	GetAllObat() []model.Obat
	GetObatByID(obatID uint64) model.Obat
	DeleteObatByID(obatID uint64) error
}

type obatConnection struct {
	connection *gorm.DB
}

func NewObatRepository(db *gorm.DB) ObatRepository {
	return &obatConnection{
		connection: db,
	}
}

func (db *obatConnection) InsertObat(obat model.Obat) model.Obat {
	db.connection.Create(&obat)
	return obat
}

func (db *obatConnection) UpdateObat(obat model.Obat) model.Obat {
	db.connection.Save(&obat)
	return obat
}

func (db *obatConnection) GetAllObat() []model.Obat {
	var obats []model.Obat
	db.connection.Find(&obats)
	return obats
}

func (db *obatConnection) GetObatByID(obatID uint64) model.Obat {
	var obat model.Obat
	db.connection.First(&obat, obatID)
	return obat
}

func (db *obatConnection) DeleteObatByID(obatID uint64) error {
	var obat model.Obat
	obat.ID = obatID
	result := db.connection.Delete(&obat)
	return result.Error
}
