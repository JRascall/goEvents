package events

// IEventEmitter is the interface that wraps the basic event emitter implementation.
type IEventEmitter interface {
	On(name string, handler func(args IEventArgs))
	Call(name string, args IEventArgs)
	Delete(name string)
}
