package hooks

// IHookSys is the interface that wraps the basic event emitter implementation.
type IHookSys interface {
	On(name string, handler func(args IHookArgs))
	Call(name string, args IHookArgs)
	Delete(name string)
}
