package helpers

import (
	"fmt"
	"gochicoba/models"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

func ScheduleCronUser(funcJob func(), params ...interface{}) {
	s := gocron.NewScheduler(time.Local)

	// parsing handled by https://pkg.go.dev/github.com/robfig/cron/v3
	// which follows https://en.wikipedia.org/wiki/Cron
	_, err := s.Cron("*/10 * * * *").Do(funcJob, params...) // every minute
	if err != nil {
		log.Println(err)
	}
	s.StartAsync()
}

func ChangeUser(db *gorm.DB) {
	var users []models.User
	err := db.Where("now() > created_at + '10 minutes'::interval").Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	for i := range users {
		err := db.Model(&users[i]).Update("name", users[i].Name+" OK").Error
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Println(users)
}
