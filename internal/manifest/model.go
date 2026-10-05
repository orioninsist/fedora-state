package manifest

type Entry struct {
	Identity string `json:"identity"`

	Name string `json:"name"`

	Type string `json:"type"`

	Version string `json:"version,omitempty"`

	Source string `json:"source,omitempty"`
}

type Manifest struct {
	Entries []Entry `json:"entries"`
}
