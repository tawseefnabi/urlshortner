package repository

import (
	"log"

	"github.com/google/uuid"
	model "github.com/tawseefnabi/urlshortner/Model"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	// Migrate the schema
	db.AutoMigrate(&model.TinyUrlData{})
	return &Repository{
		Db: db,
	}
}

func (repo *Repository) Save(urlModel model.UrlModel, hash string) {
	var urlData model.TinyUrlData
	repo.Db.Where("hash= (?)", hash).Find(&urlData)
	if urlData.Url == "" {
		repo.Db.Create(&model.TinyUrlData{
			Hash: hash,
			Url:  urlData.Url,
		})
		log.Println("Data is created for url: ", urlModel.Url, " with hash: ", hash)
	} else {
		id := uuid.New()
		repo.Db.Create(&model.TinyUrlData{
			Hash: id.String(),
			Url:  urlModel.Url,
		})
		log.Println("hash already existed, creating another hash", id.String())
	}

}
