package dao

import "time"

type ScrapeTarget struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`

	Target   string `gorm:"type:varchar(200);not null" json:"target"`
	Labels   JSON   `sql:"type:json" json:"type"`
	GroupId  int    `gorm:"type:integer;not null" json:"group_id"`
	Describe string `gorm:"type:varchar(200)" json:"describe"`
}

func (self *ScrapeTarget) TableName() string {
	return "scrape_target"
}
