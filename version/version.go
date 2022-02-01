package version

import (
	"fmt"
    "os"
    "path/filepath"
)

// All of these should be injected at build time

var applicationName string
var shortDescription string
var longDescription string

var applicationVersion string
var buildDate string

var gitVersion string

var buildOS string
var buildArch string
var goVersion string

func ExecutableName() string {
    executable, err := os.Executable()
    if err == nil {
		executable = filepath.Base(executable)
	} else {
		executable = applicationName
	}

	return executable
}

func ApplicationName() string {
	return applicationName
}

func ShortDescription() string {
	return shortDescription
}

func LongDescription() string {
	return longDescription
}

func MainVersionLine() string {
	return fmt.Sprintf("%s version %s, built %s", applicationName, applicationVersion, buildDate)
}

func GitVersionLine() string {
	return fmt.Sprintf("git version %s", gitVersion)
}

func BuildVersionLine() string {
	return fmt.Sprintf("built on %s/%s using go version %s", buildOS, buildArch, goVersion)
}

func VersionTemplate() string {
    return fmt.Sprintf("%s\n%s\n%s\n", MainVersionLine(), GitVersionLine(), BuildVersionLine())
}