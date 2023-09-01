package dao

import (
	"time"
)

type AlertRule struct {
	ID            uint          `gorm:"primary_key"`
	CreatedAt     time.Time     `gorm:comment:"创建时间"`
	UpdatedAt     time.Time     `gorm:comment:"更新时间"`
	Group         string        `json:"group,omitempty;comment:"规则分组"`
	Alert         string        `json:"alert,omitempty;comment:"规则名称"`
	Expr          string        `json:"expr,omitempty";comment:"规则表达式"`
	For           time.Duration `json:"for,omitempty;comment:"规则检测周期"`
	KeepFiringFor time.Duration `json:"keep_firing_for,omitempty;comment:"规则持续时间"`
	Labels        JSON          `sql:"type:json" json:"labels,omitempty";comment:"规则标签"`
	Annotations   JSON          `sql:"type:json" json:"annotations,omitempty";comment:"规则注解"`
}

func (self *AlertRule) TableName() string {
	return "alert_rule"
}
