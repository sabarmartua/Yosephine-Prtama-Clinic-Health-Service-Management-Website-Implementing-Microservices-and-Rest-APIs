package repository

import (
	"github.com/sabarmartua/Ulasan/model"
	"gorm.io/gorm"
)

type UlasanRepository interface {
	InsertUlasan(ulasan model.Ulasan) model.Ulasan
	UpdateUlasan(ulasan model.Ulasan) model.Ulasan
	GetAllUlasan() []model.Ulasan
	DeleteUlasan(ulasan model.Ulasan)
}

type ulasanConnection struct {
	connection *gorm.DB
}

func NewUlasanRepository(db *gorm.DB) UlasanRepository {
	return &ulasanConnection{
		connection: db,
	}
}

func (db *ulasanConnection) InsertUlasan(ulasan model.Ulasan) model.Ulasan {
	db.connection.Create(&ulasan)
	return ulasan
}

func (db *ulasanConnection) UpdateUlasan(ulasan model.Ulasan) model.Ulasan {
	db.connection.Save(&ulasan)
	return ulasan
}

func (db *ulasanConnection) GetAllUlasan() []model.Ulasan {
    var ulasan []model.Ulasan
    db.connection.Find(&ulasan)
    return ulasan
}

func (db *ulasanConnection) DeleteUlasan(ulasan model.Ulasan) {
	db.connection.Delete(&ulasan)
}
