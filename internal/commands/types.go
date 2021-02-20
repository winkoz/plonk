package commands

// OrchestratorCommand interface for executing commands against the orchestrator cli tool
type OrchestratorCommand interface {
	Deploy(manifestPath string) error
	Destroy(env string) error
	Diff(manifestPath string) error
	GetPods(namespace string) ([]byte, error)
	GetLogs(namespace string, component *string) ([]byte, error)
	Restart(namespace string, deploymentName string) ([]byte, error)
}

// BuilderCommand interface for executing commands against the builder cli tool
type BuilderCommand interface {
	Build(namespace string) error
}
