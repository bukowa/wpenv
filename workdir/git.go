package src

import (
	"os"
	"os/exec"
	"strings"
)

type Repository struct {
	*WorkDir
	origin string
}

func NewRepository(origin string, wd *WorkDir) (*Repository, error) {
	return &Repository{
		WorkDir: wd,
		origin:  origin,
	}, nil
}

func NewTempRepository(origin string) (*Repository, error) {
	r := &Repository{origin: origin}
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}
	r.WorkDir, err = NewWorkDir(tempDir)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (repo *Repository) Commands(args ...[]string) (cmd []*exec.Cmd) {
	for _, arg := range args {
		cmd = append(cmd, repo.Cmd(arg...))
	}
	return
}

func (repo *Repository) Init() *exec.Cmd {
	return repo.Cmd("init")
}

func (repo *Repository) RemoteAddOrigin() *exec.Cmd {
	return repo.Cmd("remote", "add", "origin", repo.origin)
}

func (repo *Repository) Reset(sourceBranch string) *exec.Cmd {
	return repo.Cmd("reset", strings.Join([]string{"origin", sourceBranch}, "/"))
}

func (repo *Repository) Cmd(args ...string) *exec.Cmd {
	cmd := exec.Command("git", "-C", repo.Abs())
	cmd.Dir = repo.Abs()
	cmd.Args = append(cmd.Args, args...)
	return cmd
}
