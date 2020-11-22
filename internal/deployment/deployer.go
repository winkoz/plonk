package deployment

import (
	"fmt"
	"path/filepath"

	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
	"github.com/winkoz/plonk/internal/scaffolding"
)

// Deployer creates a manifest file from templates and executes it with the deploy command.
type Deployer interface {
	Execute(ctx config.Context, stackName string) error
}

type deployer struct {
	ctx                 config.Context
	varReader           io.VariableReader
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
		templateReader: scaffolding.NewTemplateReader(ctx),
		ioService:      io.NewService(),
		templateParser: io.NewTemplateParser(),
		//TODO: change this
		orchestratorCommand: commands.NewOrchestrator("kubectl"),
	}
}

func (d deployer) Execute(ctx config.Context, env string) (err error) {
	signal := log.StartTrace("Execute")
	defer log.StopTrace(signal, err)

	log.Debugf("Ctx: \n%+v", ctx)

	// load variables
	variables, err := d.varReader.GetVariablesFromFile(ctx.ProjectName, env)
	log.Debugf("Loaded variables: %v", variables)

	// join file
	templates, err := d.environmentTemplates(env)
	log.Debugf("Loaded templates: %v", templates)
	if err != nil {
		log.Errorf("Unable to load templates for environment: %s. %v", env, err)
		return err
	}

	mainDeployFile, err := d.manifestMerger(templates, variables, env)
	if err != nil {
		log.Errorf("Unable to join all manifest files. %v", err)
		return err
	}

	deployFilePath := filepath.Join(d.ctx.TargetPath, d.ctx.DeployFolderName, fmt.Sprintf("deploy.%s", io.YAMLExtension))
	err = d.ioService.Write(deployFilePath, mainDeployFile)
	if err != nil {
		log.Errorf("Cannot save main deploy file. %+v", err)
		return err
	}

	// execute in kubectl
	err = d.orchestratorCommand.Deploy(env, deployFilePath, d.ctx)
	if err != nil {
		log.Errorf("Cannot execute deploy command %s. error = %+v", d.ctx.DeployCommand, err)
		return err
	}

	// TODO: Cleanup deploy file after success

	return nil
}

func (d deployer) environmentTemplates(env string) ([]scaffolding.TemplateData, error) {
	templateNames := []string{}
	if desiredEnv := d.ctx.Environments[env]; desiredEnv != nil {
		templateNames = desiredEnv
	} else {
		templateNames = d.ctx.Environments[baseEnvironmentKey]
	}
	log.Debugf("Loaded templates for '%s': %v", env, templateNames)

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

func (d deployer) manifestMerger(templates []scaffolding.TemplateData, deployVariables io.DeployVariables, env string) (string, error) {
	result := ""
	substitutionVariables := map[string]interface{}{}
	for key, value := range deployVariables.Build {
		substitutionVariables[key] = value
	}
	substitutionVariables[environmentVariablesKey] = deployVariables.Environment
	substitutionVariables[environmentKey] = env
	substitutionVariables[projectNameKey] = d.ctx.ProjectName

	for _, template := range templates {
		templateVariables := map[string]interface{}{}
		for key, value := range template.DefaultVariables.Build {
			if _, exists := substitutionVariables[key]; !exists {
				templateVariables[key] = value
			}
		}

		mergedVariables := io.MergeMap(templateVariables, substitutionVariables)
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
