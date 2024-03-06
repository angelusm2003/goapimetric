package models

import "time"

type Monitor_detail struct {
	DeviceID     uint      `json:"device_id"`
	Metric1      uint      `json:"metric_1"`
	Metric1Value float64   `json:"metric_1_value"`
	Metric2      uint      `json:"metric_2"`
	Metric2Value float64   `json:"metric_2_value"`
	Metric3      uint      `json:"metric_3"`
	Metric3Value float64   `json:"metric_3_value"`
	Timestamp    time.Time `json:"timestamp"`
}
