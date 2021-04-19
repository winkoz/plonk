package management

import (
	"fmt"

	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
	"github.com/winkoz/plonk/internal/io/render"
)

const allDeploymentsFlag = "--all"

// Manager allows interaction with orchestrators environments
type Manager interface {
	GetPods(env string) ([]byte, error)
	GetLogs(env string, component *string) ([]byte, error)
	Restart(ctx config.Context, env string, allDeployments bool) ([]byte, error)
}

// NewManager creates a manager object
func NewManager(ctx config.Context) Manager {
	return manager{
		ctx:                 ctx,
		orchestratorCommand: commands.NewOrchestrator(ctx, "kubectl"),
		renderer:            render.NewPlainOutputRenderer(),
	}
}

type manager struct {
	ctx                 config.Context
	orchestratorCommand commands.OrchestratorCommand
	renderer            io.OutputRenderer
}

// GetPods returns all the pods for the namespace under project-name and environment
func (m manager) GetPods(env string) (output []byte, err error) {
	signal := log.StartTrace("GetPods")
	defer log.StopTrace(signal, err)

	namespace := m.buildNamespace(env)
	output, err = m.orchestratorCommand.GetPods(namespace)
	if err != nil {
		log.Errorf("Unable to get the pods from the orchestrator. err = %v", err)
		return
	}
	m.renderer.RenderComponents(output)

	return
}

// GetLogs returns the logs for the specified `component` running on the passedin `env`
func (m manager) GetLogs(env string, component *string) (output []byte, err error) {
	signal := log.StartTrace("GetLogs")
	defer log.StopTrace(signal, err)

	namespace := m.buildNamespace(env)
	output, err = m.orchestratorCommand.GetLogs(namespace, component)
	if err != nil {
		log.Errorf("Unable to get the logs from the orchestrator. err = %v", err)
		return
	}
	m.renderer.RenderLogs(output)

	return
}

// Restart rollout restarts the deployment for the passed in `env`
func (m manager) Restart(ctx config.Context, env string, allDeployments bool) (output []byte, err error) {
	signal := log.StartTrace("Restart")
	defer log.StopTrace(signal, err)

	namespace := m.buildNamespace(env)
	deploymentName := allDeploymentsFlag
	if !allDeployments {
		deploymentName = m.buildDeploymentName(env)
	}
	output, err = m.orchestratorCommand.Restart(namespace, deploymentName)
	if err != nil {
		log.Errorf("Unable to restart the deployment from the orchestrator. err = %v", err)
		return
	}

	return
}

// *************************************************************************************
// Private methods
// *************************************************************************************

func (m manager) buildNamespace(env string) string {
	return fmt.Sprintf("%s-%s", m.ctx.ProjectName, env)
}

func (m manager) buildDeploymentName(env string) string {
	return fmt.Sprintf("%s-%s-deployment", m.ctx.ProjectName, env)
}
