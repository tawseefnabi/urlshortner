package service

import (
	model "github.com/tawseefnabi/urlshortner/Model"
	repository "github.com/tawseefnabi/urlshortner/Repository"
	utility "github.com/tawseefnabi/urlshortner/Utility"
)

var (
	UrlContent = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	UrlAddress = "localhost:8080/"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GenerateTinyUrl(urlModel model.UrlModel) model.UrlModel {
	hash := utility.ComputeHash(urlModel.Url)
	len := len(UrlContent)
	hashUrl := ""
	for hash > 0 {
		idx := hash % int64(len)
		hash = hash / int64(len)
		hashUrl += string(UrlContent[idx])
	}
	computedUrl := UrlAddress + hashUrl
	s.repo.Save(urlModel, hashUrl)
	return model.UrlModel{
		Url: computedUrl,
	}
}
