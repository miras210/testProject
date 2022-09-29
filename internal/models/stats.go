package models

type Stats struct {
	Date   CustomTime `json:"date"`
	Views  int        `json:"views,omitempty"`
	Clicks int        `json:"clicks,omitempty"`
	Cost   float32    `json:"cost,omitempty"`
}
