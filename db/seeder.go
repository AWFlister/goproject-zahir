package seeder

import (
	"encoding/json"
	"goproject-zahir/models"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Seeder struct {
	DB *gorm.DB
}

func (s Seeder) Seed() {
	path := "db/seed.json"
	bytes, _ := os.ReadFile(path)
	data_raw := make([]map[string]string, 0)
	json.Unmarshal(bytes, &data_raw)

	contacts := make([]models.Contact, 0)

	for _, v := range data_raw {
		new_con := models.Contact{
			UUID:   uuid.NewString(),
			Name:   v["name"],
			Gender: v["gender"],
			Phone:  v["phone"],
			Email:  v["email"],
		}
		contacts = append(contacts, new_con)
	}

	s.DB.Create(&contacts)
}
