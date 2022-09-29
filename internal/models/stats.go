package models

type Stats struct {
	Date   CustomTime `json:"date"`
	Views  int        `json:"views,omitempty"`
	Clicks int        `json:"clicks,omitempty"`
	Cost   float32    `json:"cost,omitempty"`
}

type FullStats struct {
	Date   CustomTime `json:"date"`
	Views  int        `json:"views,omitempty"`
	Clicks int        `json:"clicks,omitempty"`
	Cost   float32    `json:"cost,omitempty"`
	Cpc    float32    `json:"cpc"`
	Cpm    float32    `json:"cpm"`
}
