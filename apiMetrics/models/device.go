package models

import "time"

type Device struct {
	ID           uint      `json:"device_id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	MetricID1    uint      `json:"metric_id_1"`
	MetricID2    uint      `json:"metric_id_2"`
	MetricID3    uint      `json:"metric_id_3"`
	IP           string    `json:"ip"`
	DateCreation time.Time `json:"date_creation"`
}
