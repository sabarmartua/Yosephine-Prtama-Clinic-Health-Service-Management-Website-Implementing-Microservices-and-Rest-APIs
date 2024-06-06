package service

import (
	"log"

	"github.com/sabarmartua/Antrian/dto"
	"github.com/sabarmartua/Antrian/model"
	"github.com/sabarmartua/Antrian/repository"
	"github.com/mashingan/smapping"
)

type AntrianService interface {
	Create(antrianDTO dto.NewAntrianDTO) model.Antrian
	GetAll() []model.Antrian
	GetByUserID(userID uint64) []model.Antrian
	Delete(antrian model.Antrian)
}

type antrianService struct {
	antrianRepository repository.AntrianRepository
}

func NewAntrianService(antrianRepository repository.AntrianRepository) AntrianService {
	return &antrianService{
		antrianRepository: antrianRepository,
	}
}

func (service *antrianService) Create(antrianDTO dto.NewAntrianDTO) model.Antrian {
	antrian := model.Antrian{}
	err := smapping.FillStruct(&antrian, smapping.MapFields(&antrianDTO))
	if err != nil {
		log.Fatalf("Failed to map struct: %v", err)
	}

	res := service.antrianRepository.InsertAntrian(antrian)
	return res
}

func (service *antrianService) GetAll() []model.Antrian {
	return service.antrianRepository.GetAllAntrian()
}

func (service *antrianService) GetByUserID(userID uint64) []model.Antrian {
	return service.antrianRepository.GetAntrianByUserID(userID)
}

func (service *antrianService) Delete(antrian model.Antrian) {
	service.antrianRepository.DeleteAntrian(antrian)
}
