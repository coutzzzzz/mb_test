package domain

import (
	"time"
)

type Request struct {
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
	Range int       `json:"range"`
}

type Response struct {
	Timestamp int64   `json:"timestamp"`
	Mms       float64 `json:"mms"`
}

type MMS struct {
	Pair      string    `gorm:"type:varchar(6);default:null"`
	Timestamp time.Time `gorm:"default:null"`
	Mms20     float64   `gorm:"default:0"`
	Mms50     float64   `gorm:"default:0"`
	Mms200    float64   `gorm:"default:0"`
}
