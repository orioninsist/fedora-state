package model

type Package struct {
	Name string `json:"name"`

	Version string `json:"version,omitempty"`

	Release string `json:"release,omitempty"`

	Arch string `json:"arch,omitempty"`

	Manager string `json:"manager"`

	Source string `json:"source,omitempty"`
}
