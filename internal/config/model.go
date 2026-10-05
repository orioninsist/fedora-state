package config

type Config struct {
	Snapshot SnapshotConfig `yaml:"snapshot"`

	Restore RestoreConfig `yaml:"restore"`

	Providers ProviderConfig `yaml:"providers"`
}

type SnapshotConfig struct {
	Directory string `yaml:"directory"`
}

type RestoreConfig struct {
	RequireConfirmation bool `yaml:"require_confirmation"`
}

type ProviderConfig struct {
	RPM   bool `yaml:"rpm"`
	Cargo bool `yaml:"cargo"`
	UV    bool `yaml:"uv"`
}
