package model

type System struct {
	OS string `json:"os"`

	Architecture string `json:"architecture"`

	Metadata SystemMetadata `json:"metadata"`

	Objects []Object `json:"objects"`

	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`
}
