package wordpress

import (
	"github.com/pkg/errors"
	"os"
)

type Kind interface {
	New() Kind
	Name() string
	Fetch(source string) error
}

var (
	ErrorNotDir = errors.New("source is not a directory")
)

type KindLocalDirectory struct {
	source string
}

func (k *KindLocalDirectory) New() Kind {
	return &KindLocalDirectory{}
}

func (k *KindLocalDirectory) Name() string {
	return "localDir"
}

func (k *KindLocalDirectory) Fetch(source string) error {
	f, err := os.Stat(source)
	if err != nil {
		return err
	}
	if !f.IsDir() {
		return errors.Wrap(ErrorNotDir, source)
	}
	k.source = source
	return nil
}
