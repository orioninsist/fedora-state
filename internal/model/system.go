package model

import "time"

type SystemState struct {
	GeneratedAt time.Time `json:"generated_at"`

	OS string `json:"os"`

	Sources []string `json:"sources"`

	Packages []Package `json:"packages"`
}
