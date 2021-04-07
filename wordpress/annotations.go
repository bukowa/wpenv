package wordpress

import (
	"strings"
)

const (
	AnnotationPackageSource    = "wordpress.package.source"
	AnnotationPackageGroup     = "wordpress.package.group"
	AnnotationPackageKind      = "wordpress.package.kind"
	AnnotationPackageName      = "wordpress.package.name"
	AnnotationPackageVersion   = "wordpress.package.version"
	AnnotationPackageNamespace = "wordpress.package.namespace"
	AnnotationPackageURL       = "wordpress.package.url"
)

const (
	AnnotationEnvName    = "wordpress.env.name"
	AnnotationEnvThemes  = "wordpress.env.themes"
	AnnotationEnvPlugins = "wordpress.env.plugins"
)

// PackageAnnotations returns annotations for Package
func PackageAnnotations(pack *Package) map[string]string {
	return map[string]string{
		AnnotationPackageSource:    pack.Source,
		AnnotationPackageGroup:     string(pack.Group),
		AnnotationPackageKind:      string(pack.Kind),
		AnnotationPackageName:      pack.Name,
		AnnotationPackageVersion:   pack.Version,
		AnnotationPackageNamespace: pack.Namespace,
		AnnotationPackageURL:       pack.URL,
	}
}

// EnvAnnotations returns annotations for Env
func EnvAnnotations(env *Env) map[string]string {
	return map[string]string{
		AnnotationEnvName:    env.Name,
		AnnotationEnvThemes:  strings.Join(env.Themes, ","),
		AnnotationEnvPlugins: strings.Join(env.Plugins, ","),
	}
}
