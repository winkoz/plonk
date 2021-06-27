package events

type CommandSelected interface {
	Type() MessageType
	Data() interface{}
	GetCommandName() string
}

type commandSelectedData struct {
	CommandName string
}

type commandSelectedMessage struct {
	data commandSelectedData
}

func NewCommandSelectedMessage(commandName string) CommandSelected {
	return commandSelectedMessage{
		data: commandSelectedData{
			CommandName: commandName,
		},
	}
}

func (c commandSelectedMessage) Type() MessageType {
	return CommandSelectedMessage
}

func (c commandSelectedMessage) Data() interface{} {
	return c.data
}

func (c commandSelectedMessage) GetCommandName() string {
	return c.data.CommandName
}
