package wordpress

type (
	// Config represents configuration of Wordpress Envs for a given Repository.
	Config struct {
		Kind string         `yaml:"kind"` // 'theme' or 'plugin'
		Envs map[string]Env `yaml:"envs,omitempty" json:",omitempty"`
	}

	// Env represents specific Wordpress environment
	Env struct {
		Name    string   `yaml:"-"`
		Package *Package `yaml:"-"`

		Themes  []string `yaml:",omitempty" json:",omitempty"`
		Plugins []string `yaml:",omitempty" json:",omitempty"`
		Scripts []string `yaml:",omitempty" json:",omitempty"`
	}
)
