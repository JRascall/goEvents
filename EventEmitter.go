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

func (e *eventEmitter) On(name string, handler func(args IEventArgs)) {
	lowerName := strings.ToLower(name)
	e.listeners[lowerName] = append(e.listeners[lowerName], handler)
}

func (e *eventEmitter) Call(name string, args IEventArgs) {
	lowerName := strings.ToLower(name)
	if e.listeners[lowerName] != nil {
		for _, handler := range e.listeners[lowerName] {
			handler(args)
		}
	}
}

func (e *eventEmitter) Delete(name string) {
	lowerName := strings.ToLower(name)
	if e.listeners[lowerName] != nil {
		delete(e.listeners, lowerName)
	}
}
