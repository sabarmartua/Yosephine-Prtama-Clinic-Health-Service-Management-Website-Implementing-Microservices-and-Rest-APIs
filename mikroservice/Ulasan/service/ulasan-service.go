package service

import (
	"log"

	"github.com/sabarmartua/Ulasan/dto"
	"github.com/sabarmartua/Ulasan/model"
	"github.com/sabarmartua/Ulasan/repository"
	"github.com/mashingan/smapping"
)

type UlasanService interface {
	Create(ulasanDTO dto.NewUlasanDTO) model.Ulasan
	Update(ulasanDTO dto.UpdateUlasanDTO) model.Ulasan
	Delete(ulasan model.Ulasan)
	GetAll() []model.Ulasan
}

type ulasanService struct {
	ulasanRepository repository.UlasanRepository
}

func NewUlasanService(ulasanRepository repository.UlasanRepository) UlasanService {
	return &ulasanService{
		ulasanRepository: ulasanRepository,
	}
}

func (service *ulasanService) Create(ulasanDTO dto.NewUlasanDTO) model.Ulasan {
	ulasan := model.Ulasan{}
	err := smapping.FillStruct(&ulasan, smapping.MapFields(&ulasanDTO))
	if err != nil {
		log.Fatalf("Failed to map struct: %v", err)
	}

	res := service.ulasanRepository.InsertUlasan(ulasan)
	return res
}

func (service *ulasanService) Update(ulasanDTO dto.UpdateUlasanDTO) model.Ulasan {
	ulasan := model.Ulasan{}
	err := smapping.FillStruct(&ulasan, smapping.MapFields(&ulasanDTO))
	if err != nil {
		log.Fatalf("Failed to map struct: %v", err)
	}

	res := service.ulasanRepository.UpdateUlasan(ulasan)
	return res
}

func (service *ulasanService) Delete(ulasan model.Ulasan) {
	service.ulasanRepository.DeleteUlasan(ulasan)
}

func (service *ulasanService) GetAll() []model.Ulasan {
	return service.ulasanRepository.GetAllUlasan()
}
