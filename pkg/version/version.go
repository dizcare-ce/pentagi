// Package version provides build-time version information for pentagi.
package version

import (
	"fmt"
	"runtime"
)

// Build information, populated at build time via ldflags.
var (
	// Version is the semantic version of the application.
	Version = "dev"

	// GitCommit is the git commit hash at build time.
	GitCommit = "none"

	// GitBranch is the git branch at build time.
	GitBranch = "unknown"

	// BuildDate is the date the binary was built.
	BuildDate = "unknown"

	// GoVersion is the version of Go used to compile.
	GoVersion = runtime.Version()
)

// Info holds all version-related metadata.
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	GitBranch string `json:"git_branch"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

// Get returns the current version information.
func Get() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		GitBranch: GitBranch,
		BuildDate: BuildDate,
		GoVersion: GoVersion,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String returns a human-readable version string.
// Example output:
//
//	pentagi version 1.2.3 (commit: abc1234, branch: main, built: 2024-01-01, go1.22.0, linux/amd64)
func (i Info) String() string {
	return fmt.Sprintf(
		"pentagi version %s (commit: %s, branch: %s, built: %s, %s, %s)",
		i.Version,
		i.GitCommit,
		i.GitBranch,
		i.BuildDate,
		i.GoVersion,
		i.Platform,
	)
}
