package main

import (
	"fmt"
)

// name is the name of the runtime
const name = "kata-runtime"

// name of the project
const project = "kata-containers"

// prefix used to denote non-standard CLI commands and options.
const projectPrefix = "kata"

// original URL for this project
const projectURL = "https://github.com/kata-containers"

const defaultRootDirectory = "/var/run/kata"

// commit is the git commit the runtime is compiled from.
// sourced from ldflags
var commit string

// version is the runtime version.
// sourced from ldflags
var version string

// project-specific command names
var envCmd = fmt.Sprintf("%s-env", projectPrefix)
var checkCmd = fmt.Sprintf("%s-check", projectPrefix)

// project-specific option names
var configFilePathOption = fmt.Sprintf("%s-config", projectPrefix)
var showConfigPathsOption = fmt.Sprintf("%s-show-default-config-paths", projectPrefix)

// Default config file used by stateless systems.
var defaultRuntimeConfiguration = "/usr/share/kata-containers/configuration.toml"

// Alternate config file that takes precedence over
// defaultRuntimeConfiguration.
var defaultSysConfRuntimeConfiguration = "/etc/kata-containers/configuration.toml"
