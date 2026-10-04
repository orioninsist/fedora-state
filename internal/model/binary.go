package model

type Binary struct {
	Name string `json:"name"`

	Path string `json:"path"`

	RealPath string `json:"real_path,omitempty"`

	Owner string `json:"owner,omitempty"`
}
