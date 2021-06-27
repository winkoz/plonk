package events

type Message interface {
	Type() MessageType
	Data() interface{}
}
