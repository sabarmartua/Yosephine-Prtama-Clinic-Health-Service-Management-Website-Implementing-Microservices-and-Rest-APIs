package service

import (
	
    "github.com/sabarmartua/FAQ/dto"
    "github.com/sabarmartua/FAQ/model"
    "github.com/sabarmartua/FAQ/repository"
)

type FAQService interface {
    Insert(b dto.NewFAQDTO) (model.FAQ, error)
    Update(b dto.UpdateFAQDTO) (model.FAQ, error)
    Delete(id uint64) error
    All() []model.FAQ
    FindByID(id uint64) model.FAQ
}

type faqService struct {
    faqRepository repository.FAQRepository
}

func NewFAQService(faqRepository repository.FAQRepository) FAQService {
    return &faqService{
        faqRepository: faqRepository,
    }
}

func (service *faqService) All() []model.FAQ {
    return service.faqRepository.All()
}

func (service *faqService) FindByID(id uint64) model.FAQ {
    return service.faqRepository.FindByID(id)
}

func (service *faqService) Insert(b dto.NewFAQDTO) (model.FAQ, error) {
    faq := model.FAQ{
        Pertanyaan: b.Pertanyaan,
        Jawaban:    b.Jawaban,
    }

    res := service.faqRepository.InsertFAQ(faq)
    return res, nil
}

func (service *faqService) Update(b dto.UpdateFAQDTO) (model.FAQ, error) {
    faq := model.FAQ{
        ID:         b.ID,
        Pertanyaan: b.Pertanyaan,
        Jawaban:    b.Jawaban,
    }

    res := service.faqRepository.UpdateFAQ(faq)
    return res, nil
}

func (service *faqService) Delete(id uint64) error {
    return service.faqRepository.DeleteByID(id)
}
