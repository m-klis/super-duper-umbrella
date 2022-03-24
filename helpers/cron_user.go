package helpers

import (
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
	// fmt.Println("Change User Runned")
	var users []models.User
	err := db.Where("checked = ?", false).
		Where("(created_at + '10 minutes'::interval) < now()").
		Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	for i := range users {
		err := db.Model(&users[i]).
			Updates(map[string]interface{}{"name": users[i].Name + " OK", "checked": true}).
			Error
		if err != nil {
			log.Println(err)
		}
	}
	// fmt.Println(users)
}
