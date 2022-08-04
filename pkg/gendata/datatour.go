package gendata

import (
	"github.com/google/uuid"
	"github.com/luispfcanales/sse-go/internal/core/domain"
)

//GenerateTours setup tours
func GenerateTours(tours map[string]*domain.Tour) {
	id := getNewID()
	pictures := generatePictures()
	tours[id] = &domain.Tour{
		ID:          id,
		CompanyName: "Inkaterra",
		Description: "Es una agencia de turismo",
		Photo:       "/static/tours/image-slide-1.jpg",
		Pictures:    pictures,
	}
	id = getNewID()
	tours[id] = &domain.Tour{
		ID:          id,
		CompanyName: "Explorer-In",
		Description: "Es una agencia de turismo",
		Photo:       "/static/tours/image-slide-2.jpg",
		Pictures:    pictures,
	}
	id = getNewID()
	tours[id] = &domain.Tour{
		ID:          id,
		CompanyName: "Wasai",
		Description: "Es una agencia de turismo",
		Photo:       "/static/tours/image-slide-3.jpg",
		Pictures:    pictures,
	}
	id = getNewID()
	tours[id] = &domain.Tour{
		ID:          id,
		CompanyName: "Tambopata Expedition",
		Description: "Es una agencia de turismo",
		Photo:       "/static/tours/image-slide-4.jpg",
		Pictures:    pictures,
	}
}

//getNewId
func getNewID() string {
	return uuid.New().String()
}
func generatePictures() []domain.Picture {
	return []domain.Picture{
		{ID: getNewID(), URL: "/static/tours/image-slide-1.jpg"},
		{ID: getNewID(), URL: "/static/tours/image-slide-2.jpg"},
		{ID: getNewID(), URL: "/static/tours/image-slide-3.jpg"},
		{ID: getNewID(), URL: "/static/tours/image-slide-4.jpg"},
		{ID: getNewID(), URL: "/static/tours/image-slide-5.jpg"},
		{ID: getNewID(), URL: "/static/tours/image-slide-6.jpg"},
		{ID: getNewID(), URL: "/static/tours/image-slide-7.jpg"},
		{ID: getNewID(), URL: "/static/tours/image-slide-8.jpg"},
		{ID: getNewID(), URL: "/static/tours/image-slide-9.jpg"},
	}
}
