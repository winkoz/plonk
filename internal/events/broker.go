package events

type Broker interface {
	PostMessage(msg Message)
	GetBrokerChannel() <-chan Message
}

type broker struct {
	bus chan Message
}

func NewBroker() Broker {
	return &broker{
		bus: make(chan Message, 1),
	}
}

func (b broker) PostMessage(msg Message) {
	b.bus <- msg
}

func (b broker) GetBrokerChannel() <-chan Message {
	return b.bus
}
