package service

import (
    "log"

	"github.com/sabarmartua/Kategori-Artikel/dto"
	"github.com/sabarmartua/Kategori-Artikel/model"
	"github.com/sabarmartua/Kategori-Artikel/repository"
    "github.com/mashingan/smapping"
)

// KategoriArtikelService adalah kontrak tentang apa yang dapat dilakukan oleh layanan ini
type KategoriArtikelService interface {
    Insert(b dto.NewKategoriArtikelDTO) model.KategoriArtikel
    Update(b dto.UpdateKategoriArtikelDTO) model.KategoriArtikel
    Delete(id uint64) error
    All() []model.KategoriArtikel
    FindByID(id uint64) model.KategoriArtikel
}

type kategoriArtikelService struct {
    kategoriArtikelRepository repository.KategoriArtikelRepository
}

// NewKategoriArtikelService membuat instance baru dari KategoriArtikelService
func NewKategoriArtikelService(kategoriArtikelRepository repository.KategoriArtikelRepository) KategoriArtikelService {
    return &kategoriArtikelService{
        kategoriArtikelRepository: kategoriArtikelRepository,
    }
}

func (service *kategoriArtikelService) All() []model.KategoriArtikel {
    return service.kategoriArtikelRepository.All()
}

func (service *kategoriArtikelService) FindByID(id uint64) model.KategoriArtikel {
    return service.kategoriArtikelRepository.FindByID(id)
}


func (service *kategoriArtikelService) Insert(b dto.NewKategoriArtikelDTO) model.KategoriArtikel {
    kategoriArtikel := model.KategoriArtikel{}
    err := smapping.FillStruct(&kategoriArtikel, smapping.MapFields(&b))
    if err != nil {
        log.Fatalf("Gagal memetakan DTO ke model: %v", err)
    }

    res := service.kategoriArtikelRepository.InsertKategoriArtikel(kategoriArtikel)
    return res
}

func (service *kategoriArtikelService) Update(b dto.UpdateKategoriArtikelDTO) model.KategoriArtikel {
    kategoriArtikel := model.KategoriArtikel{}
    err := smapping.FillStruct(&kategoriArtikel, smapping.MapFields(&b))
    if err != nil {
        log.Fatalf("Gagal memetakan DTO ke model: %v", err)
    }

    res := service.kategoriArtikelRepository.UpdateKategoriArtikel(kategoriArtikel)
    return res
}

func (service *kategoriArtikelService) Delete(id uint64) error {
    return service.kategoriArtikelRepository.DeleteByID(id)
}