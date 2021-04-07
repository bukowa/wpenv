package wordpress

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	// WpOrgGroup means source from wordpress.org
	WpOrgGroup Group = "wporg"
	// GitGroup means source from git repositories
	GitGroup Group = "git"

	// ThemeKind means it's a theme
	ThemeKind Kind = "theme"
	// PluginKind means it's a plugin
	PluginKind Kind = "plugin"
	// ScriptKind means it's a custom shell script.
	ScriptKind Kind = "script"

	// GitVersionDefault is a default version if nothing was provided
	GitVersionDefault = "master"
)

type (
	// Group represents group of Package object
	// for ex. it may be a git source (GitGroup) or wordpress.org source (WpOrgGroup)
	Group string

	// Kind represents kind of Package objects
	// for ex. it may be theme (ThemeKind) or a plugin (PluginKind)
	Kind string

	// Package represents versioned object
	Package struct {
		Source string // full text passed "http://gitlab.com/group/project@master" "woocommerce@4.2.3"

		Group Group // ex. WpOrgGroup - determined automatically from passed Source
		Kind  Kind  // ex. ThemeKind, PluginKind - has to be set via SetKind

		Name      string // first part before @
		Version   string // second part ex. "master" or "1.0.0" - can be empty for WpOrgGroup
		Namespace string // repository namespace ex. "group/project"

		Host string // for ex. gitlab.com, empty for other group than GitGroup

		// Attributes set by SetKind
		URL      string // https://github.com/$NAMESPACE/tree/$VERSION | https://wordpress.org/plugins/$NAME
		FetchURL string // https://downloads.wordpress.org/$NAME.$VERSION.zip | https://gitlab.com/$NAMESPACE@$VERSION
	}
)

// NewPackage creates new Package objects
func NewPackage(source string) (*Package, error) {
	p, v := getVersion(source)
	ver := &Package{
		Source:  source,
		Name:    p,
		Version: v,
	}

	u, err := url.Parse(source)
	if err != nil {
		return nil, err
	}

	// Empty scheme - still try to determine if its a url scheme for ex. gitlab.com/asd/asd
	if !u.IsAbs() {
		if hasS(source, ".") && hasS(source, "/") {
			ld := lastI(source, ".")
			fs := firstI(source, "/")
			u.Host = source[:ld] + source[ld:fs]
			u.Scheme = "https"
		}
	}

	// Quit early
	if u.Host == "" {
		ver.Group = WpOrgGroup
		return ver, nil
	}

	// Otherwise handle Git repository
	ver.Group = GitGroup

	// Set default version
	if ver.Version == "" {
		ver.Version = GitVersionDefault
	}

	// GetResource namespace ex. group/project
	if ver.Namespace, err = parseNamespace(ver.Name); err != nil {
		return nil, err
	}

	// Name should not contains URI scheme
	if u.Scheme != "" {
		ver.Name = strings.Replace(ver.Name, u.Scheme+"://", "", 1)
	}
	ver.Host = u.Host
	return ver, nil
}

// SetKind set's Kind and automatically builds URI's for known Kind's
func (pkg *Package) SetKind(kind Kind) {
	pkg.Kind = kind
	switch pkg.Group {
	case GitGroup:
		pkg.URL = fmt.Sprintf("https://%s/%s/tree/%s", pkg.Host, pkg.Namespace, pkg.Version)
		pkg.FetchURL = fmt.Sprintf("https://%s/%s@%s", pkg.Host, pkg.Namespace, pkg.Version)
	case WpOrgGroup:
		pkg.FetchURL = buildWpOrgURL(pkg.Name, pkg.Version, string(pkg.Kind))
		switch kind {
		case ThemeKind:
			pkg.URL = fmt.Sprintf("https://wordpress.org/themes/%s", pkg.Name)
		case PluginKind:
			pkg.URL = fmt.Sprintf("https://wordpress.org/plugins/%s", pkg.Name)
		}
	}
}

func parseNamespace(uri string) (string, error) {
	var u *url.URL
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	// Handle "@"
	if i := strings.Index(u.Path, "@"); i != -1 {
		u.Path = u.Path[:i]
	}
	s := strings.Split(u.Path, "/")
	if len(s) < 2 {
		return "", fmt.Errorf("error: parsing project namespace for uri: %s", uri)
	}
	return strings.Join([]string{s[1], s[2]}, "/"), nil
}

func getVersion(text string) (p, v string) {
	s := strings.Split(text, "@")
	if len(s) >= 2 {
		p = s[0]
		v = strings.Join(s[1:], "@")
	} else {
		p = text
		v = ""
	}
	return
}

func buildWpOrgURL(pkg, version, kind string) string {
	uri := fmt.Sprintf("https://downloads.wordpress.org/%s", kind)
	if version == "" {
		return fmt.Sprintf("%s/%s.zip", uri, pkg)
	}
	return fmt.Sprintf("%s/%s.%s.zip", uri, pkg, version)
}

func hasS(s, sub string) bool {
	return strings.Contains(s, sub)
}

func lastI(s, sub string) int {
	return strings.LastIndex(s, sub)
}

func firstI(s, sub string) int {
	return strings.Index(s, sub)
}
