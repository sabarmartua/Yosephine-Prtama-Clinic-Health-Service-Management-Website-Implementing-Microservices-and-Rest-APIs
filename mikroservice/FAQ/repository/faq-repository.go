package repository

import (
    "github.com/sabarmartua/FAQ/model"
    "gorm.io/gorm"
)

type FAQRepository interface {
    InsertFAQ(faq model.FAQ) model.FAQ
    UpdateFAQ(faq model.FAQ) model.FAQ
    All() []model.FAQ
    FindByID(faqID uint64) model.FAQ
    DeleteByID(faqID uint64) error
}

type faqConnection struct {
    connection *gorm.DB
}

func NewFAQRepository(db *gorm.DB) FAQRepository {
    return &faqConnection{
        connection: db,
    }
}

func (db *faqConnection) InsertFAQ(faq model.FAQ) model.FAQ {
    db.connection.Create(&faq)
    return faq
}

func (db *faqConnection) UpdateFAQ(faq model.FAQ) model.FAQ {
    db.connection.Save(&faq)
    return faq
}

func (db *faqConnection) All() []model.FAQ {
    var faqs []model.FAQ
    db.connection.Find(&faqs)
    return faqs
}

func (db *faqConnection) FindByID(faqID uint64) model.FAQ {
    var faq model.FAQ
    db.connection.First(&faq, faqID)
    return faq
}

func (db *faqConnection) DeleteByID(faqID uint64) error {
    var faq model.FAQ
    faq.ID = faqID
    result := db.connection.Delete(&faq)
    return result.Error
}
