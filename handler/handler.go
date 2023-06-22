package handler

import (
	"goproject-zahir/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Migrate() {
	models := []models.Model{&models.Contact{}}
	for _, model := range models {
		h.DB.AutoMigrate(model)
	}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/", h.HomeIndex)

	r.Route("/contacts", func(r chi.Router) {
		r.Get("/", h.ContactsIndex)
		r.Post("/", h.ContactsCreate)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.ContactsShow)
			r.Put("/", h.ContactsUpdate)
			r.Delete("/", h.ContactsDestroy)
		})
	})
}
