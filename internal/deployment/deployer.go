package deployment

import (
	"fmt"
	"path/filepath"

	"github.com/winkoz/plonk/internal"
	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
	"github.com/winkoz/plonk/internal/scaffolding"
)

// Deployer creates a manifest file from templates and executes it with the deploy command.
type Deployer interface {
	Execute(ctx config.Context, stackName string, tagName string, dryRun bool) error
}

type deployer struct {
	ctx                 config.Context
	varReader           io.VariableReader
	secretsReader       io.SecretReader
	templateReader      scaffolding.TemplateReader
	ioService           io.Service
	templateParser      io.TemplateParser
	orchestratorCommand commands.OrchestratorCommand
}

// NewDeployer creates a deployer object
func NewDeployer(ctx config.Context) Deployer {
	return deployer{
		ctx:            ctx,
		varReader:      io.NewVariableReader(filepath.Join(ctx.TargetPath, ctx.DeployVariablesPath)),
		secretsReader:  io.NewSecretReader(filepath.Join(ctx.TargetPath, ctx.DeploySecretsPath)),
		templateReader: scaffolding.NewTemplateReader(ctx),
		ioService:      io.NewService(),
		templateParser: io.NewTemplateParser(),
		//TODO: https://github.com/winkoz/plonk/issues/59 change this
		orchestratorCommand: commands.NewOrchestrator(ctx, "kubectl"),
	}
}

func (d deployer) Execute(ctx config.Context, env string, tagName string, dryRun bool) (err error) {
	signal := log.StartTrace("Execute")
	defer log.StopTrace(signal, err)

	log.Debugf("Ctx: \n%+v", ctx)

	// load variables
	variables, err := d.varReader.GetVariablesFromFile(ctx.ProjectName, env)
	log.Debugf("Loaded variables: %v", variables)

	if tagName != "" {
		variables.Build["DOCKER_IMAGE"] = tagName
		variables.Environment["DOCKER_IMAGE"] = tagName
	}

	// load secrets
	secrets, err := d.secretsReader.GetSecretsFromFile(ctx.ProjectName, env)
	log.Debugf("Loaded secrets: %v", secrets)

	// join file
	templates, err := d.environmentTemplates(env)
	log.Debugf("Loaded templates: %v", templates)
	if err != nil {
		log.Errorf("Unable to load templates for environment: %s. %v", env, err)
		return err
	}

	mainDeployFile, err := d.manifestMerger(templates, variables, secrets, env)
	if err != nil {
		log.Errorf("Unable to join all manifest files. %v", err)
		return err
	}

	deployFilePath := filepath.Join(d.ctx.DeployFolderName, fmt.Sprintf("deploy.%s", io.YAMLExtension))
	deployFullPath := filepath.Join(d.ctx.TargetPath, deployFilePath)
	err = d.ioService.Write(deployFullPath, mainDeployFile)
	if err != nil {
		log.Errorf("Cannot save main deploy file. %+v", err)
		return err
	}

	// execute in kubectl
	cmd := d.orchestratorCommand.Deploy
	if dryRun {
		cmd = d.orchestratorCommand.Diff
	}

	err = cmd(deployFilePath)
	if err != nil {
		log.Errorf("Cannot execute deploy command %s. error = %+v", d.ctx.DeployCommand, err)
		return err
	}

	if !dryRun { // Delete the deploy.yaml
		d.ioService.DeletePath(deployFullPath)
	}

	return nil
}

// *************************************************************************************
// Private methods
// *************************************************************************************

func (d deployer) environmentTemplates(env string) ([]scaffolding.TemplateData, error) {
	templateNames := []string{}
	if desiredEnv := d.ctx.Environments[env]; desiredEnv != nil {
		templateNames = desiredEnv
	} else {
		templateNames = d.ctx.Environments[internal.BaseEnvironmentKey]
	}
	log.Debugf("Loaded templates for '%s': %v", env, templateNames)

	//inject namespace template
	templateNames = append([]string{"namespace"}, templateNames...)

	result := make([]scaffolding.TemplateData, len(templateNames))
	for _, templateName := range templateNames {
		log.Debugf("Loading template: %s", templateName)
		template, err := d.templateReader.Read(templateName)
		if err != nil {
			return nil, err
		}
		result = append(result, template)
	}

	return result, nil
}

func (d deployer) manifestMerger(templates []scaffolding.TemplateData, deployVariables io.DeployVariables, deploySecrets io.DeploySecrets, env string) (string, error) {
	result := ""
	substitutionVariables := map[string]interface{}{}
	substitutionEnvironmentVariables := deployVariables.Environment

	for key, value := range deployVariables.Build {
		substitutionVariables[key] = value
	}

	substitutionVariables[environmentKey] = env
	substitutionVariables[projectNameKey] = d.ctx.ProjectName

	for _, template := range templates {
		templateVariables := map[string]interface{}{}
		templateEnvVariables := map[string]string{}
		for key, value := range template.DefaultVariables.Build {
			if _, exists := substitutionVariables[key]; !exists {
				templateVariables[key] = value
			}
		}

		for key, value := range template.DefaultVariables.Environment {
			if _, exists := substitutionVariables[key]; !exists {
				templateEnvVariables[key] = value
			}
		}

		mergedEnvVariables := io.MergeStringMap(templateEnvVariables, substitutionEnvironmentVariables)
		mergedVariables := io.MergeMap(templateVariables, substitutionVariables)
		mergedVariables[environmentVariablesKey] = mergedEnvVariables
		mergedVariables[environmentSecretsKey] = deploySecrets.Secret
		for _, manifestFileName := range template.Manifests {
			data, err := d.ioService.ReadFile(manifestFileName)
			if err != nil {
				return "", err
			}
			contents := string(data)

			parsedTemplate, err := d.templateParser.Parse(mergedVariables, contents)
			if err != nil {
				return "", err
			}
			result += parsedTemplate
			result += "\n---\n"
		}
	}

	return result, nil
}
