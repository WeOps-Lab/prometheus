package dao

import "time"

type ScrapeGroup struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`

	Name         string `gorm:"type:varchar(200);not null" json:"name"`
	MetricsPath  string `gorm:"type:varchar(200);not null" json:"metrics_path"`
	Labels       JSON   `sql:"type:json" json:"labels"`
	Scheme       string `gorm:"type:varchar(200);not null" json:"scheme"`
	Interval     int    `gorm:"type:integer;not null" json:"interval"`
	Timeout      int    `gorm:"type:integer;not null" json:"timeout"`
	Describe     string `gorm:"type:varchar(200)" json:"describe"`
	AuthUser     string `gorm:"type:varchar(200)" json:"auth_user"`
	AuthPassword string `gorm:"type:varchar(200)" json:"auth_password"`
	AuthToken    string `gorm:"type:varchar(200)" json:"auth_token"`
}

func (self *ScrapeGroup) TableName() string {
	return "scrape_group"
}
