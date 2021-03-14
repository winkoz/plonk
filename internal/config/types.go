package config

import "path/filepath"

// Constants
var defaultCustomTemplatesPath = filepath.Join("$HOME", ".plonk", "templates")

const deployFolderName = "deploy"

var deployVariablesPath = filepath.Join(deployFolderName, "variables")

var deploySecretsPath = filepath.Join(deployFolderName, "secrets")

const deployDeployCommand = "kubectl"

const deployBuildCommand = "docker"

const registryDefaultValue = "registry.hub.docker.com"
