package management

import (
	"fmt"

	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

// Manager allows interaction with orchestrators environments
type Manager interface {
	GetPods(env string) ([]byte, error)
	GetLogs(env string) ([]byte, error)
}

// NewManager creates a manager object
func NewManager(ctx config.Context) Manager {
	return manager{
		ctx:                 ctx,
		orchestratorCommand: commands.NewOrchestrator(ctx, "kubectl"),
		renderer:            io.NewPlainOutputRenderer(),
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

func (m manager) GetLogs(env string) (output []byte, err error) {
	signal := log.StartTrace("GetLogs")
	defer log.StopTrace(signal, err)

	namespace := m.buildNamespace(env)
	output, err = m.orchestratorCommand.GetLogs(namespace)

	if err != nil {
		log.Errorf("Unable to get the logs from the orchestrator. err = %v", err)
		return
	}

	m.renderer.RenderComponents(output)

	return
}

func (m manager) buildNamespace(env string) string {
	return fmt.Sprintf("%s-%s", m.ctx.ProjectName, env)
}
