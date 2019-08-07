package postgresql

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	// Import postgresql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/l-lin/8-phone/phone"
)

const phoneNumberTable = "phone_numbers"

// Repository that access to a PostgreSQL database
type Repository struct {
	dbURL string
}

// New postgresql repository
func New(dbURL string) *Repository {
	return &Repository{dbURL}
}

// GetAll phones from PostgreSQL database
func (r *Repository) GetAll() []*phone.Phone {
	db, err := gorm.Open("postgres", r.dbURL)
	if err != nil {
		log.WithField("err", err).Fatal("Could not open db")
	}
	defer db.Close()
	var phones []*phone.Phone
	if err := db.Table(phoneNumberTable).Find(&phones).Error; err != nil {
		log.WithField("err", err).Fatal("Could not fetch phones from db")
	}
	return phones
}
