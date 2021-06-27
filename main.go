package main

import (
	"fmt"
	"sync"

	"github.com/winkoz/plonk/cmd"
	"github.com/winkoz/plonk/internal"
	"github.com/winkoz/plonk/internal/events"
)

func main() {
	var wg sync.WaitGroup

	runtimeCtx := internal.RuntimeContext{
		Broker: events.NewBroker(),
	}

	wg.Add(1)
	go func(ctx internal.RuntimeContext) {
		cmd.Execute(ctx)
	}(runtimeCtx)

	wg.Add(1)
	go func(ctx internal.RuntimeContext) {
		for m := range ctx.Broker.GetBrokerChannel() {
			fmt.Printf("Message %+v Received.", m)
		}
	}(runtimeCtx)

	wg.Wait()
}
