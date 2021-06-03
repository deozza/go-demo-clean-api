package entity

import (
	"time"
)

type Status string

const (
	DRAFT Status = "DRAFT"
	ONGOING     = "ONGOING"
	ENDED       = "ENDED"
)

type Party struct {
	ID           ID
	Name         string
	Description  string
	Host         ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Status       Status
	Participants []ID
	Objects      []ID
}

func (p *Party) ChangeStatus(status Status) {
	p.Status = status
}