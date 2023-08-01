package dto

import (
	"time"
)

type GetResponseV1Payload struct {
	Type       string     `json:"type"`
	Contact    Contact    `json:"contact"`
	Account    Account    `json:"account"`
	Message    Message    `json:"message,omitempty"`
	ClickTrack ClickTrack `json:"clickTrack,omitempty"`
	SMS        SMS        `json:"sms,omitempty"`
	Event      Event      `json:"event"`
}

type Event struct {
	OccurredAt time.Time `json:"occurredAt"`
}
