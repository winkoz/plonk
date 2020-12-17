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
	output, err = m.executeCommand("GetPods", m.orchestratorCommand.GetPods, env)

	return
}

func (m manager) GetLogs(env string) (output []byte, err error) {
	output, err = m.executeCommand("GetLogs", m.orchestratorCommand.GetLogs, env)

	return
}

//-----------------------------
// Private Methods
//-----------------------------

func (m manager) buildNamespace(env string) string {
	return fmt.Sprintf("%s-%s", m.ctx.ProjectName, env)
}

func (m manager) executeCommand(logName string, command func(string) ([]byte, error), env string) (output []byte, err error) {
	signal := log.StartTrace(logName)
	defer log.StopTrace(signal, err)

	namespace := m.buildNamespace(env)
	output, err = command(namespace)

	if err != nil {
		log.Errorf("Unable to execute command %s in orchestrator. err = %v", logName, err)
		return
	}

	m.renderer.RenderComponents(output)

	return
}
