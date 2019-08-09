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

// Count the number of phones that have the given phone number
func (r *Repository) Count(value string) int {
	db, err := gorm.Open("postgres", r.dbURL)
	if err != nil {
		log.WithField("err", err).Fatal("Could not open db")
	}
	defer db.Close()
	var p []*phone.Phone
	if err := db.Table(phoneNumberTable).Where("value = ?", value).Find(&p).Error; err != nil {
		log.WithFields(log.Fields{
			"err":   err,
			"value": value,
		}).Fatal("Could not fetch phone from db")
	}
	if p != nil {
		return len(p)
	}
	return 0
}

// Update a phone in the db
func (r *Repository) Update(p *phone.Phone) {
	db, err := gorm.Open("postgres", r.dbURL)
	if err != nil {
		log.WithField("err", err).Fatal("Could not open db")
	}
	defer db.Close()
	if err := db.Table(phoneNumberTable).Save(&p).Error; err != nil {
		log.WithFields(log.Fields{
			"err":   err,
			"phone": *p,
		}).Fatal("Could not update phone from db")
	}
}

// Delete a phone in the db by its id
func (r *Repository) Delete(id int) {
	p := &phone.Phone{ID: id}
	db, err := gorm.Open("postgres", r.dbURL)
	if err != nil {
		log.WithField("err", err).Fatal("Could not open db")
	}
	defer db.Close()
	if err := db.Table(phoneNumberTable).Delete(&p).Error; err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Fatal("Could not delete phone from db")
	}
}
