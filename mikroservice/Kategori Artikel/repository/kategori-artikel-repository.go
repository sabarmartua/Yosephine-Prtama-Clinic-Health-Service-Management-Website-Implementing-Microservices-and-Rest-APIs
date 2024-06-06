package repository

import (
    "github.com/sabarmartua/Kategori-Artikel/model"
    "gorm.io/gorm"
)

type KategoriArtikelRepository interface {
    InsertKategoriArtikel(kategoriArtikel model.KategoriArtikel) model.KategoriArtikel
    UpdateKategoriArtikel(kategoriArtikel model.KategoriArtikel) model.KategoriArtikel
    All() []model.KategoriArtikel
    FindByID(kategoriArtikelID uint64) model.KategoriArtikel
    DeleteByID(kategoriArtikelID uint64) error
}

type kategoriArtikelConnection struct {
    connection *gorm.DB
}

func NewKategoriArtikelRepository(db *gorm.DB) KategoriArtikelRepository {
    return &kategoriArtikelConnection{
        connection: db,
    }
}

func (db *kategoriArtikelConnection) InsertKategoriArtikel(kategoriArtikel model.KategoriArtikel) model.KategoriArtikel {
    db.connection.Create(&kategoriArtikel)
    return kategoriArtikel
}

func (db *kategoriArtikelConnection) UpdateKategoriArtikel(kategoriArtikel model.KategoriArtikel) model.KategoriArtikel {
    db.connection.Save(&kategoriArtikel)
    return kategoriArtikel
}

func (db *kategoriArtikelConnection) All() []model.KategoriArtikel {
    var kategoriArtikels []model.KategoriArtikel
    db.connection.Find(&kategoriArtikels)
    return kategoriArtikels
}

func (db *kategoriArtikelConnection) FindByID(kategoriArtikelID uint64) model.KategoriArtikel {
    var kategoriArtikel model.KategoriArtikel
    db.connection.First(&kategoriArtikel, kategoriArtikelID)
    return kategoriArtikel
}

func (db *kategoriArtikelConnection) DeleteByID(kategoriArtikelID uint64) error {
    var kategoriArtikel model.KategoriArtikel
    kategoriArtikel.ID = kategoriArtikelID
    result := db.connection.Delete(&kategoriArtikel)
    return result.Error
}
