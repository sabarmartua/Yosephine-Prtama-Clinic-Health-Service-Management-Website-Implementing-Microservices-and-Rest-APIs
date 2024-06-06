package repository

import (
    "github.com/sabarmartua/CutiDokter/model"
    "gorm.io/gorm"
)

type CutiDokterRepository interface {
    InsertCutiDokter(cutiDokter model.CutiDokter) model.CutiDokter
    UpdateCutiDokter(cutiDokter model.CutiDokter) model.CutiDokter
    All() []model.CutiDokter
    FindByID(cutiDokterID uint64) model.CutiDokter
    DeleteByID(cutiDokterID uint64) error
}

type cutiDokterConnection struct {
    connection *gorm.DB
}

func NewCutiDokterRepository(db *gorm.DB) CutiDokterRepository {
    return &cutiDokterConnection{
        connection: db,
    }
}

func (db *cutiDokterConnection) InsertCutiDokter(cutiDokter model.CutiDokter) model.CutiDokter {
    db.connection.Create(&cutiDokter)
    return cutiDokter
}

func (db *cutiDokterConnection) UpdateCutiDokter(cutiDokter model.CutiDokter) model.CutiDokter {
    db.connection.Save(&cutiDokter)
    return cutiDokter
}

func (db *cutiDokterConnection) All() []model.CutiDokter {
    var cutiDokters []model.CutiDokter
    db.connection.Find(&cutiDokters)
    return cutiDokters
}

func (db *cutiDokterConnection) FindByID(cutiDokterID uint64) model.CutiDokter {
    var cutiDokter model.CutiDokter
    db.connection.First(&cutiDokter, cutiDokterID)
    return cutiDokter
}

func (db *cutiDokterConnection) DeleteByID(cutiDokterID uint64) error {
    var cutiDokter model.CutiDokter
    cutiDokter.ID = cutiDokterID
    result := db.connection.Delete(&cutiDokter)
    return result.Error
}
