package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/sabarmartua/Artikel/dto"
	"github.com/sabarmartua/Artikel/model"
	"github.com/sabarmartua/Artikel/repository"
)

type ArtikelService interface {
	All() ([]model.Artikel, error)
	FindByID(artikelID uint64) (model.Artikel, error)
	Create(dto dto.NewArtikelDTO) (model.Artikel, error)
	Update(artikelID uint64, dto dto.UpdateArtikelDTO) (model.Artikel, error)
	Delete(artikelID uint64) error
	GetRelatedByCategory(categoryID uint64) ([]model.Artikel, error)
}

type artikelService struct {
	artikelRepository repository.ArtikelRepository
}

func NewArtikelService(artikelRepository repository.ArtikelRepository) ArtikelService {
	return &artikelService{
		artikelRepository: artikelRepository,
	}
}

func (s *artikelService) All() ([]model.Artikel, error) {
	return s.artikelRepository.All()
}

func (s *artikelService) FindByID(artikelID uint64) (model.Artikel, error) {
	return s.artikelRepository.FindByID(artikelID)
}

func (s *artikelService) Create(dto dto.NewArtikelDTO) (model.Artikel, error) {
	artikel := model.Artikel{}

	err := smapping.FillStruct(&artikel, smapping.MapFields(&dto))
	if err != nil {
		log.Fatalf("Failed to map DTO to model: %v", err)
	}

	return s.artikelRepository.Insert(artikel)
}

func (s *artikelService) Update(artikelID uint64, dto dto.UpdateArtikelDTO) (model.Artikel, error) {
	artikel, err := s.artikelRepository.FindByID(artikelID)
	if err != nil {
		return model.Artikel{}, err
	}

	err = smapping.FillStruct(&artikel, smapping.MapFields(&dto))
	if err != nil {
		return model.Artikel{}, err
	}

	return s.artikelRepository.Update(artikelID, artikel)
}

func (s *artikelService) Delete(artikelID uint64) error {
	return s.artikelRepository.Delete(artikelID)
}

func (s *artikelService) GetRelatedByCategory(categoryID uint64) ([]model.Artikel, error) {
	return s.artikelRepository.GetRelatedByCategory(categoryID)
}
