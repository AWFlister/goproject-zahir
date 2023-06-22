package handler

import (
	"encoding/json"
	"fmt"
	"goproject-zahir/models"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ContactInput struct {
	Name   string
	Gender string
	Phone  string
	Email  string
}

type ContactList struct {
	Data      []models.Contact `json:"data"`
	Page      int              `json:"page"`
	PageTotal int              `json:"page_total"`
}

func (h *Handler) ContactsIndex(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	contacts := []models.Contact{}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	order_by := r.URL.Query().Get("order_by")
	filterstr := strings.Split(r.URL.Query().Get("filter"), ",")
	var filter [][]string
	if r.URL.Query().Has("filter") {
		for _, f := range filterstr {
			filter = append(filter, strings.Split(f, ":"))
		}
	}

	offset := limit * (page - 1)
	var count int64 = 0
	query := h.DB.Model(&models.Contact{}).Offset(offset).Limit(limit).Order(order_by)
	if r.URL.Query().Has("filter") {
		for _, f := range filter {
			switch f[0] {
			case "name":
				query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", f[1]))
			case "gender":
				query.Where("gender LIKE ?", f[1])
			case "phone":
				query.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", f[1]))
			case "email":
				query.Where("email LIKE ?", fmt.Sprintf("%%%s%%", f[1]))
			}
		}
	}
	query.Find(&contacts)
	query.Count(&count)

	w.Header().Set("Content-Type", "application/json")
	clist := ContactList{
		Data:      contacts,
		Page:      page,
		PageTotal: int(math.Ceil(float64(count) / float64(limit))),
	}
	encoder.Encode(clist)
}

func (h *Handler) ContactsCreate(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	var input ContactInput
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&input)

	contact := models.Contact{
		UUID:   uuid.NewString(),
		Name:   input.Name,
		Gender: input.Gender,
		Phone:  input.Phone,
		Email:  input.Email,
	}

	h.DB.Create(&contact)

	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(contact)
}

func (h *Handler) ContactsShow(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	id := chi.URLParam(r, "id")

	contact := models.Contact{}
	h.DB.First(&contact, "uuid = ?", id)

	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(contact)
}

func (h *Handler) ContactsUpdate(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	id := chi.URLParam(r, "id")
	contact := models.Contact{}

	h.DB.First(&contact, "uuid = ?", id)

	var input ContactInput
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&input)

	h.DB.Model(&contact).Updates(input)

	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(contact)
}

func (h *Handler) ContactsDestroy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	contact := models.Contact{}

	h.DB.Delete(&contact, "uuid = ?", id)
}
