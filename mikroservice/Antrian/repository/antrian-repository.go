package repository

import (
	"github.com/sabarmartua/Antrian/model"
	"gorm.io/gorm"
)

type AntrianRepository interface {
	InsertAntrian(antrian model.Antrian) model.Antrian
	GetAllAntrian() []model.Antrian
	GetAntrianByUserID(userID uint64) []model.Antrian
	DeleteAntrian(antrian model.Antrian)
}

type antrianConnection struct {
	connection *gorm.DB
}

func NewAntrianRepository(db *gorm.DB) AntrianRepository {
	return &antrianConnection{
		connection: db,
	}
}

func (db *antrianConnection) InsertAntrian(antrian model.Antrian) model.Antrian {
	db.connection.Create(&antrian)
	return antrian
}

func (db *antrianConnection) GetAllAntrian() []model.Antrian {
	var antrian []model.Antrian
	db.connection.Find(&antrian)
	return antrian
}

func (db *antrianConnection) GetAntrianByUserID(userID uint64) []model.Antrian {
	var antrian []model.Antrian
	db.connection.Where("user_id = ?", userID).Find(&antrian)
	return antrian
}

func (db *antrianConnection) DeleteAntrian(antrian model.Antrian) {
	db.connection.Delete(&antrian)
}
