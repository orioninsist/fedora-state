package model

type Package struct {
	Name string `json:"name"`

	Version string `json:"version,omitempty"`

	Release string `json:"release,omitempty"`

	Arch string `json:"arch,omitempty"`

	Manager string `json:"manager"`

	Source string `json:"source,omitempty"`

	Repository string `json:"repository,omitempty"`

	InstallMethod string `json:"install_method,omitempty"`

	BinaryPath string `json:"binary_path,omitempty"`

	Verified bool `json:"verified"`
}
