package model

type Source struct {
	Name      string `json:"name"`
	Available bool   `json:"available"`
	Collected bool   `json:"collected"`
}
