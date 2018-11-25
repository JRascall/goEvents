package events

type eventArgs struct {
	data map[string]interface{}
}

// CreateEventArgs creates a new event args to be passed to an event handler
func CreateEventArgs(data map[string]interface{}) IEventArgs {
	return &eventArgs{
		data: data,
	}
}

func (e *eventArgs) Data() map[string]interface{} {
	return e.data
}
