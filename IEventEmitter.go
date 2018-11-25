package events

type IEventEmitter interface {
	On(name string, handler func(args IEventArgs))
	Call(name string, args IEventArgs)
	Delete(name string)
}
