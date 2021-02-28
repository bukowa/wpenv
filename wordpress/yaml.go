package wordpress

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type (
	// Config represents configuration of wordpress environments.
	Config struct {
		Envs map[string]*Env `yaml:"envs,omitempty" json:",omitempty"`
	}

    // Package represents a plugin/theme or other kind of files that needs to be installed.
    Package struct {
        // Type determines way of fetching the files.
        Type string `yaml:"kind" json:"kind"`
        // Source is the location of files.
        Source string `yaml:"source" json:"source"`
        // Kind determines interface used to fetch files.
        Kind Kind
    }

	// Env represents specific wordpress environment
	Env struct {
		Name    string   `yaml:"name" json:"name"`
		Packages []*Package `yaml:"packages" json:"packages"`
	}
)

var (
	ErrorTypeNotHandled = errors.New("type is not handled")
)

func (c *Config) Parse(kinds []Kind) error {
	for _, e := range c.Envs {
		for _, p := range e.Packages {
			var parsed bool
			for _, k := range kinds {
				if k.Name() == p.Type {
					p.Kind = k.New()
					err := p.Kind.Fetch(p.Source)
					parsed = true
					if err != nil {
						return errors.Wrapf(err, "while fetching %s by kind %s", p.Source, k.Name())
					}
					break
				}
			}
			if !parsed {
				return errors.Wrap(ErrorTypeNotHandled, p.Type)
			}
		}
	}
	return nil
}

func NewConfigFromReader(r io.Reader) (*Config, error) {
	return configFromReader(r)
}

func configFromReader(r io.Reader) (*Config, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
