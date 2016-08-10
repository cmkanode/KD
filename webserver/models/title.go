package models

import "time"

type (
	// Title represents the structure of our resource
	Title struct {
		TitleId              int       `json:"TitleId" bson:"TitleId"`
		TitleName            string    `json:"TitleName" bson:"TitleName"`
		TitleNameSortable    string    `json:"TitleNameSortable" bson:"TitleNameSortable"`
		TitleTypeId          int       `json:"TitleTypeId" bson:"TitleTypeId"`
		ReleaseYear          int       `json:"ReleaseYear" bson:"ReleaseYear"`
		ProcessedDateTimeUTC time.Time `json:"ProcessedDateTimeUTC" bson:"ProcessedDateTimeUTC"`
	}
)
