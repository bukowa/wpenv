package wordpress

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

func Test_getVersion(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name  string
		args  args
		wantP string
		wantV string
	}{
		{name: "1", args: args{text: "https://gitlab.com/wp-env/example-plugin@example@v0.1.0"}, wantP: "https://gitlab.com/wp-env/example-plugin", wantV: "example@v0.1.0"},
		{name: "2", args: args{text: "gitlab.com/wp-env/example-plugin@"}, wantP: "gitlab.com/wp-env/example-plugin", wantV: ""},
		{name: "2", args: args{text: "gitlab.com/wp-env/example-plugin@a"}, wantP: "gitlab.com/wp-env/example-plugin", wantV: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotV := getVersion(tt.args.text)
			if gotP != tt.wantP {
				t.Errorf("getVersion() gotP = %v, want %v", gotP, tt.wantP)
			}
			if gotV != tt.wantV {
				t.Errorf("getVersion() gotV = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func Test_parseNamespace(t *testing.T) {
	tests := map[string]string{
		"https://gitlab.com/aspengrovestudios/wp-layouts/-/merge_requests":                 "aspengrovestudios/wp-layouts",
		"https://gitlab.com/aspengrovestudios/wp-layouts/":                                 "aspengrovestudios/wp-layouts",
		"https://gitlab.com/aspengrovestudios/wp-layouts/-/":                               "aspengrovestudios/wp-layouts",
		"https://gitlab.com/aspengrovestudios/wp-layouts/-/packages":                       "aspengrovestudios/wp-layouts",
		"https://gitlab.com/aspengrovestudios/divi-switch-rebuild/-/tree/etm":              "aspengrovestudios/divi-switch-rebuild",
		"https://gitlab.com/aspengrovestudios/wp-layouts@test":                             "aspengrovestudios/wp-layouts",
		"https://gitlab.com/wp-env/divi@b65b335b8b0d6e72405se44dcbf67cbaec683f3cf":         "wp-env/divi",
		"https://gitlab.com/wp-env/divi/-/tree/b65b335b8b0d6e72405e44dcbf67cbaec683f3cf":   "wp-env/divi",
		"https://gitlab.com/wp-env/divi/-/commit/b65b335b8b0d6e72405e44dcbf67cbaec683f3cf": "wp-env/divi",
		"https://gitlab.com/wp-env/divi":                                                   "wp-env/divi",
		"https://gitlab.com/wp-env/divi/-/tree/4.5.6":                                      "wp-env/divi",
	}
	for k, v := range tests {
		namespace, err := parseNamespace(k)
		if err != nil {
			t.Error(err)
		}
		if namespace != v {
			t.Error(namespace, k, v)
		}
	}
}

func TestNewVersioned1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    *Package
		wantErr bool
	}{
		{name: "1", args: args{text: "woocommerce"}, want: &Package{
			Group:     WpOrgGroup,
			Source:    "woocommerce",
			Name:      "woocommerce",
			Version:   "",
			Namespace: "",
			Host:      "",
		}, wantErr: false},
		{name: "2", args: args{text: "woocommerce@4.3.3"}, want: &Package{
			Group:     WpOrgGroup,
			Source:    "woocommerce@4.3.3",
			Name:      "woocommerce",
			Version:   "4.3.3",
			Namespace: "",
			Host:      "",
		}, wantErr: false},
		{name: "3", args: args{text: "https://github.com/bukowa/template@branch"}, want: &Package{
			Group:     GitGroup,
			Source:    "https://github.com/bukowa/template@branch",
			Name:      "github.com/bukowa/template",
			Version:   "branch",
			Namespace: "bukowa/template",
			Host:      "github.com",
		}, wantErr: false},
		{name: "4", args: args{text: "https://github.com/bukowa/template"}, want: &Package{
			Group:     GitGroup,
			Source:    "https://github.com/bukowa/template",
			Name:      "github.com/bukowa/template",
			Version:   "master",
			Namespace: "bukowa/template",
			Host:      "github.com",
		}},
		{name: "5", args: args{text: "http://github.com/bukowa/template"}, want: &Package{
			Group:     GitGroup,
			Source:    "http://github.com/bukowa/template",
			Name:      "github.com/bukowa/template",
			Version:   "master",
			Namespace: "bukowa/template",
			Host:      "github.com",
		}},
		{name: "5", args: args{text: "github.com/bukowa/template"}, want: &Package{
			Group:     GitGroup,
			Source:    "github.com/bukowa/template",
			Name:      "github.com/bukowa/template",
			Version:   "master",
			Namespace: "bukowa/template",
			Host:      "github.com",
		}},
		{name: "6", args: args{text: "github.com/bukowa/template@tag"}, want: &Package{
			Group:     GitGroup,
			Source:    "github.com/bukowa/template@tag",
			Name:      "github.com/bukowa/template",
			Version:   "tag",
			Namespace: "bukowa/template",
			Host:      "github.com",
		}},
		{name: "7", args: args{text: "https://gitlab.com/bukowa/template@branch"}, want: &Package{
			Group:     GitGroup,
			Source:    "https://gitlab.com/bukowa/template@branch",
			Name:      "gitlab.com/bukowa/template",
			Version:   "branch",
			Namespace: "bukowa/template",
			Host:      "gitlab.com",
		}, wantErr: false},
		{name: "8", args: args{text: "https://gitlab.com/bukowa/template"}, want: &Package{
			Group:     GitGroup,
			Source:    "https://gitlab.com/bukowa/template",
			Name:      "gitlab.com/bukowa/template",
			Version:   "master",
			Namespace: "bukowa/template",
			Host:      "gitlab.com",
		}},
		{name: "9", args: args{text: "http://gitlab.com/bukowa/template"}, want: &Package{
			Group:     GitGroup,
			Source:    "http://gitlab.com/bukowa/template",
			Name:      "gitlab.com/bukowa/template",
			Version:   "master",
			Namespace: "bukowa/template",
			Host:      "gitlab.com",
		}},
		{name: "10", args: args{text: "gitlab.com/bukowa/template"}, want: &Package{
			Group:     GitGroup,
			Source:    "gitlab.com/bukowa/template",
			Name:      "gitlab.com/bukowa/template",
			Version:   "master",
			Namespace: "bukowa/template",
			Host:      "gitlab.com",
		}},
		{name: "11", args: args{text: "gitlab.com/bukowa/template@tag"}, want: &Package{
			Group:     GitGroup,
			Source:    "gitlab.com/bukowa/template@tag",
			Name:      "gitlab.com/bukowa/template",
			Version:   "tag",
			Namespace: "bukowa/template",
			Host:      "gitlab.com",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPackage(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPackage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func TestVersioned_SetKind(t *testing.T) {
	type fields struct {
		Package string
		Version string
	}
	type args struct {
		kind Kind
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "1", fields: fields{Package: "twentyfifteen", Version: "1.2.3"}, args: args{kind: ThemeKind}},
		{name: "2", fields: fields{Package: "twentyfifteen", Version: "1.2.3"}, args: args{kind: PluginKind}},
		{name: "1", fields: fields{Package: "twentyfifteen", Version: ""}, args: args{kind: ThemeKind}},
		{name: "2", fields: fields{Package: "twentyfifteen", Version: ""}, args: args{kind: PluginKind}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Package{
				Name:    tt.fields.Package,
				Version: tt.fields.Version,
			}
			v.Group = WpOrgGroup
			v.SetKind(tt.args.kind)
			var x string
			switch tt.args.kind {
			case ThemeKind:
				switch tt.fields.Version {
				default:
					x = fmt.Sprintf("https://downloads.wordpress.org/theme/%s.%s.zip", tt.fields.Package, tt.fields.Version)
				case "":
					x = fmt.Sprintf("https://downloads.wordpress.org/theme/%s.zip", tt.fields.Package)
				}
			case PluginKind:
				switch tt.fields.Version {
				default:
					x = fmt.Sprintf("https://downloads.wordpress.org/plugin/%s.%s.zip", tt.fields.Package, tt.fields.Version)
				case "":
					x = fmt.Sprintf("https://downloads.wordpress.org/plugin/%s.zip", tt.fields.Package)
				}
			default:
				t.Error("unhandled")
			}
			log.Print(x)
		})
	}
}
