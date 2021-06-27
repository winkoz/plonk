package internal

import "github.com/winkoz/plonk/internal/events"

const BaseEnvironmentKey string = "base"

type RuntimeContext struct {
	Broker events.Broker
}
