package kjspkg

import "github.com/Modern-Modpacks/kjspkg/pkg/commons"

type Config struct {
	Installed map[string][]string `json:"installed"`

	Version   int       `json:"version"`
	ModLoader ModLoader `json:"modloader"`
}

type ModLoader string

func (s ModLoader) String() string {
	return commons.TitleCase(string(s))
}

var ModLoaders = []ModLoader{MLForge, MLFabric, MLQuilt, MLNeoforge}

const (
	MLForge    ModLoader = "forge"
	MLFabric             = "fabric"
	MLQuilt              = "quilt"
	MLNeoforge           = "neoforge"
)

func DefaultConfig() *Config {
	return &Config{
		Installed: map[string][]string{},
	}
}
