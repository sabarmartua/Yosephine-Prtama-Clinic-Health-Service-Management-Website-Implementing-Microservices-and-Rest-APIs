package service

import (
    "errors"
    "log"
    "time"

    "github.com/mashingan/smapping"
    "github.com/sabarmartua/Obat/dto"
    "github.com/sabarmartua/Obat/model"
    "github.com/sabarmartua/Obat/repository"
)

type ObatService interface {
    Insert(obatDTO dto.NewObatDTO) (model.Obat, error)
    Update(obatDTO dto.UpdateObatDTO) (model.Obat, error)
    Delete(id uint64) error
    GetAll() []model.Obat
    GetByID(id uint64) model.Obat
}

type obatService struct {
    obatRepository repository.ObatRepository
}

func NewObatService(obatRepository repository.ObatRepository) ObatService {
    return &obatService{
        obatRepository: obatRepository,
    }
}

func (service *obatService) Insert(obatDTO dto.NewObatDTO) (model.Obat, error) {
    // Validating expiration date
    if obatDTO.ExpiredDate.Before(time.Now()) {
        return model.Obat{}, errors.New("Expired date must be in the future")
    }

    obat := model.Obat{}
    err := smapping.FillStruct(&obat, smapping.MapFields(&obatDTO))
    if err != nil {
        log.Fatalf("Failed to map DTO to model: %v", err)
    }

    res := service.obatRepository.InsertObat(obat)
    return res, nil
}

func (service *obatService) Update(obatDTO dto.UpdateObatDTO) (model.Obat, error) {
    // Validating expiration date
    if obatDTO.ExpiredDate.Before(time.Now()) {
        return model.Obat{}, errors.New("Expired date must be in the future")
    }

    obat := model.Obat{}
    err := smapping.FillStruct(&obat, smapping.MapFields(&obatDTO))
    if err != nil {
        log.Fatalf("Failed to map DTO to model: %v", err)
    }

    res := service.obatRepository.UpdateObat(obat)
    return res, nil
}

func (service *obatService) Delete(id uint64) error {
    return service.obatRepository.DeleteObatByID(id)
}

func (service *obatService) GetAll() []model.Obat {
    return service.obatRepository.GetAllObat()
}

func (service *obatService) GetByID(id uint64) model.Obat {
    return service.obatRepository.GetObatByID(id)
}
