package config

import "path/filepath"

// Constants
var defaultCustomTemplatesPath = filepath.Join("$HOME", ".plonk", "templates")

const deployFolderName = "deploy"

var deployVariablesPath = filepath.Join(deployFolderName, "variables")

const deployDeployCommand = "kubectl"
