package hooks

type hookArgs struct {
	data map[string]interface{}
}

// CreateEventArgs creates a new event args to be passed to an event handler
func CreateHookArgs(data map[string]interface{}) IHookArgs {
	return &hookArgs{
		data: data,
	}
}

func (e *hookArgs) Data() map[string]interface{} {
	return e.data
}
