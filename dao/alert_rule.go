package dao

import (
	"time"
)

type AlertRule struct {
	ID            string        `json:"id,omitempty"`
	Group         string        `json:"group" json:"group,omitempty"`
	Alert         string        `json:"alert" json:"alert,omitempty"`
	Expr          string        `json:"expr" json:"expr,omitempty"`
	For           time.Duration `json:"for" json:"for,omitempty"`
	KeepFiringFor time.Duration `json:"keepFiringFor" json:"keepFiringFor,omitempty"`
	Labels        JSON          `json:"labels" json:"labels,omitempty"`
	Annotations   JSON          `json:"annotations" json:"annotations,omitempty"`
}
