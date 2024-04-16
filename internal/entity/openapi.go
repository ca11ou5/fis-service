package entity

import "time"

type ScoringStatus struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	RequestID string    `json:"requestId" bson:"requestId"`
	Status    string    `json:"status" bson:"status"`
	Decision  string    `json:"decision" bson:"decision"`
	Comment   string    `json:"comment,omitempty" bson:"comment,omitempty"`
}
