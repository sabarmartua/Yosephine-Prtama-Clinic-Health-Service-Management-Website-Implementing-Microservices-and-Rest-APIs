package repository

import (
	"github.com/sabarmartua/Artikel/model"
	"gorm.io/gorm"
)

type ArtikelRepository interface {
	Insert(artikel model.Artikel) (model.Artikel, error)
	Update(artikelID uint64, artikel model.Artikel) (model.Artikel, error)
	All() ([]model.Artikel, error)
	FindByID(artikelID uint64) (model.Artikel, error)
	Delete(artikelID uint64) error
	GetRelatedByCategory(categoryID uint64) ([]model.Artikel, error)
}

type artikelRepository struct {
	db *gorm.DB
}

func NewArtikelRepository(db *gorm.DB) ArtikelRepository {
	return &artikelRepository{
		db: db,
	}
}

func (r *artikelRepository) Insert(artikel model.Artikel) (model.Artikel, error) {
	err := r.db.Create(&artikel).Error
	return artikel, err
}

func (r *artikelRepository) Update(artikelID uint64, artikel model.Artikel) (model.Artikel, error) {
	err := r.db.Model(&model.Artikel{}).Where("id = ?", artikelID).Updates(&artikel).Error
	return artikel, err
}

func (r *artikelRepository) All() ([]model.Artikel, error) {
	var artikels []model.Artikel
	err := r.db.Find(&artikels).Error
	return artikels, err
}

func (r *artikelRepository) FindByID(artikelID uint64) (model.Artikel, error) {
	var artikel model.Artikel
	err := r.db.First(&artikel, artikelID).Error
	return artikel, err
}

func (r *artikelRepository) Delete(artikelID uint64) error {
	return r.db.Delete(&model.Artikel{}, artikelID).Error
}

func (r *artikelRepository) GetRelatedByCategory(categoryID uint64) ([]model.Artikel, error) {
	var artikels []model.Artikel
	err := r.db.Where("kategori_id = ?", categoryID).Find(&artikels).Error
	return artikels, err
}
