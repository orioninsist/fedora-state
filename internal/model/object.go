package model

type Evidence struct {
	Type  string `json:"type"`
	Path  string `json:"path,omitempty"`
	Value string `json:"value,omitempty"`
}

type Object struct {
	Name string `json:"name"`

	Type string `json:"type"`

	Identity string `json:"identity,omitempty"`

	Location string `json:"location,omitempty"`

	Version string `json:"version,omitempty"`

	Source string `json:"source,omitempty"`

	Evidence []Evidence `json:"evidence,omitempty"`
}
