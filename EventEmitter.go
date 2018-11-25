package events

import (
	"strings"
)

type eventEmitter struct {
	listeners map[string][]func(IEventArgs)
}

// CreateEventEmitter creates an event emitter used handle events
func CreateEventEmitter() IEventEmitter {
	return &eventEmitter{
		listeners: make(map[string][]func(IEventArgs)),
	}
}

// On will register a callback to an eventname
func (e *eventEmitter) On(name string, handler func(args IEventArgs)) {
	lowerName := strings.ToLower(name)
	e.listeners[lowerName] = append(e.listeners[lowerName], handler)
}

// Call will invoke an event by its name with the supplied args
func (e *eventEmitter) Call(name string, args IEventArgs) {
	lowerName := strings.ToLower(name)
	if e.listeners[lowerName] != nil {
		for _, handler := range e.listeners[lowerName] {
			handler(args)
		}
	}
}

// Delete will delete all handlers of the event name supplied
func (e *eventEmitter) Delete(name string) {
	lowerName := strings.ToLower(name)
	if e.listeners[lowerName] != nil {
		delete(e.listeners, lowerName)
	}
}
