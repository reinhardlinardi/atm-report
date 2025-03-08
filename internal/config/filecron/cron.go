package filecron

type Config struct {
	Path    string `yaml:"path"`
	WaitSec uint   `yaml:"wait_sec"`
}

func Parse(config *Config) error {
	return nil
}
