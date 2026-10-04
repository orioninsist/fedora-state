package model

type SystemMetadata struct {
	Distribution string `json:"distribution,omitempty"`

	DistributionID string `json:"distribution_id,omitempty"`

	Version string `json:"version,omitempty"`

	VersionID string `json:"version_id,omitempty"`

	Kernel string `json:"kernel,omitempty"`

	Hostname string `json:"hostname,omitempty"`
}
