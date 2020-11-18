package commands

// NewOrchestrator this will return a class to execute actions on the orchestrator command line tool
func NewOrchestrator(orchestratorType string) OrchestratorCommand {
	return kubecltCommand{
		executor: NewExecutor(),
	}
}
