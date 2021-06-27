package events

type MessageType int

const (
	None MessageType = iota
	CommandSelectedMessage
	PodLoadedMessage
)
