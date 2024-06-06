package service

import (
    "log"
    "time"
    "errors"

	"github.com/sabarmartua/CutiDokter/dto"
	"github.com/sabarmartua/CutiDokter/model"
	"github.com/sabarmartua/CutiDokter/repository"
    "github.com/mashingan/smapping"
)
type CutiDokterService interface {
	Insert(b dto.NewCutiDokterDTO) (model.CutiDokter, error)
	Update(b dto.UpdateCutiDokterDTO) (model.CutiDokter, error)
	Delete(id uint64) error
	All() []model.CutiDokter
	FindByID(id uint64) model.CutiDokter
}

type cutiDokterService struct {
	cutiDokterRepository repository.CutiDokterRepository
}

func NewCutiDokterService(cutiDokterRepository repository.CutiDokterRepository) CutiDokterService {
	return &cutiDokterService{
		cutiDokterRepository: cutiDokterRepository,
	}
}

func (service *cutiDokterService) All() []model.CutiDokter {
	return service.cutiDokterRepository.All()
}

func (service *cutiDokterService) FindByID(id uint64) model.CutiDokter {
	return service.cutiDokterRepository.FindByID(id)
}

func (service *cutiDokterService) Insert(b dto.NewCutiDokterDTO) (model.CutiDokter, error) {
	// Validasi tanggal
	if b.TanggalSelesai.Before(b.TanggalMulai) {
		return model.CutiDokter{}, errors.New("Tanggal selesai tidak boleh lebih awal dari tanggal mulai")
	}

	now := time.Now()
	if b.TanggalMulai.Before(now) {
		return model.CutiDokter{}, errors.New("Tanggal mulai tidak boleh sebelum tanggal sekarang")
	}

	cutiDokter := model.CutiDokter{}
	err := smapping.FillStruct(&cutiDokter, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Gagal memetakan DTO ke model: %v", err)
	}

	res := service.cutiDokterRepository.InsertCutiDokter(cutiDokter)
	return res, nil
}

func (service *cutiDokterService) Update(b dto.UpdateCutiDokterDTO) (model.CutiDokter, error) {
	// Validasi tanggal
	if b.TanggalSelesai.Before(b.TanggalMulai) {
		return model.CutiDokter{}, errors.New("Tanggal selesai tidak boleh lebih awal dari tanggal mulai")
	}

	now := time.Now()
	if b.TanggalMulai.Before(now) {
		return model.CutiDokter{}, errors.New("Tanggal mulai tidak boleh sebelum tanggal sekarang")
	}

	cutiDokter := model.CutiDokter{}
	err := smapping.FillStruct(&cutiDokter, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Gagal memetakan DTO ke model: %v", err)
	}

	res := service.cutiDokterRepository.UpdateCutiDokter(cutiDokter)
	return res, nil
}

func (service *cutiDokterService) Delete(id uint64) error {
	return service.cutiDokterRepository.DeleteByID(id)
}
