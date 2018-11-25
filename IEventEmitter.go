package events

// IEventEmitter is the interface that wraps the basic event emitter implementation.
//
// On(name string, handlerFunc func(args IEventArgs)) will register a callback to an eventname
//
// Call(name string, args IEventArgs) will invoke an event by its name with the supplied args
//
// Delete(name string) will delete all handlers of the event name supplied
type IEventEmitter interface {
	On(name string, handler func(args IEventArgs))
	Call(name string, args IEventArgs)
	Delete(name string)
}
