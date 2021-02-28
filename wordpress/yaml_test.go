package wordpress_test

import (
	"errors"
	"os"
	"testing"
)
import . "github.com/bukowa/wpenv/wordpress"

func TestConfig_Parse(t *testing.T) {
	type fields struct {
		Envs map[string]*Env
	}
	type args struct {
		kinds []Kind
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err error
	}{
		{
			name: "localDir valid",
			fields: struct{ Envs map[string]*Env }{Envs: map[string]*Env{
				"env1": {
					Name:     "test1",
					Packages: []*Package{
						{
							Type:   "localDir",
							Source: ".",
						},
					},
				},
			}},
			args: struct{ kinds []Kind }{kinds:
				[]Kind{&KindLocalDirectory{}},
			},
			wantErr: false,
		},
		{
			name: "not handled Type on Package",
			fields: struct{ Envs map[string]*Env }{Envs: map[string]*Env{
				"env1": {
					Name:     "test1",
					Packages: []*Package{
						{
							Type:   "unhandled",
							Source: ".",
						},
					},
				},
			}},
			args: struct{ kinds []Kind }{kinds:
			[]Kind{&KindLocalDirectory{}},
			},
			wantErr: true,
			err: ErrorTypeNotHandled,
		},
		{
			name: "localDir with file",
			fields: struct{ Envs map[string]*Env }{Envs: map[string]*Env{
				"env1": {
					Name:     "test1",
					Packages: []*Package{
						{
							Type:   "localDir",
							Source: "yaml_test.go",
						},
					},
				},
			}},
			args: struct{ kinds []Kind }{kinds:
			[]Kind{&KindLocalDirectory{}},
			},
			wantErr: true,
			err: ErrorNotDir,
		},
		{
			name: "localDir with file does not exists",
			fields: struct{ Envs map[string]*Env }{Envs: map[string]*Env{
				"env1": {
					Name:     "test1",
					Packages: []*Package{
						{
							Type:   "localDir",
							Source: "thisfileshouldneverbecreated123.unkonwnext",
						},
					},
				},
			}},
			args: struct{ kinds []Kind }{kinds:
			[]Kind{&KindLocalDirectory{}},
			},
			wantErr: true,
			err: os.ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Envs: tt.fields.Envs,
			}
			err := c.Parse(tt.args.kinds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = '%v', wantErr '%v'", err, tt.wantErr)
			}
			if err != nil {
				if !errors.Is(err, tt.err) {
					t.Errorf("Parse() error ='%v', wantErr '%v'", err, tt.err)
				}
			}
		})
	}
}
