package models

type (
	// Award represents our award resource
	Award struct {
		Award        string `json:"Award" bson:"Award"`
		AwardCompany string `json:"AwardCompany" bson:"AwardCompany"`
	}
)
