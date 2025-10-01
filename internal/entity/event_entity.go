package entity

import (
	"fmt"
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Event struct {
	Base
	Slug     string    `gorm:"type:varchar(255);not null" json:"slug"`
	Title    string    `gorm:"type:varchar(255);not null" json:"title"`
	Price    uint64    `gorm:"not null" json:"price"`
	Location string    `gorm:"type:varchar(255);not null" json:"location"`
	Date     time.Time `gorm:"not null" json:"date"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	if e.Slug == "" {
		e.Slug = slug.Make(e.Title)
		var count int64
		for {
			err = tx.Model(&Event{}).Where("slug = ?", e.Slug).Count(&count).Error
			if err != nil {
				return err
			}

			if count == 0 {
				break
			}

			e.Slug = fmt.Sprintf("%s-%d", slug.Make(e.Title), time.Now().UnixNano())
		}
	}

	return nil
}
