package hooks

import (
	"strings"
)

type hookSystem struct {
	hooks map[string][]func(IHookArgs)
}

// CreateHookSystem creates an hook emitter used handle events
func CreateHookSystem() IHookSys {
	return &hookSystem{
		hooks: make(map[string][]func(IHookArgs)),
	}
}

// On will register a callback to an hookname
func (h *hookSystem) On(name string, handler func(args IHookArgs)) {
	lowerName := strings.ToLower(name)
	h.hooks[lowerName] = append(h.hooks[lowerName], handler)
}

// Call will invoke an hook by its name with the supplied args
func (h *hookSystem) Call(name string, args IHookArgs) {
	lowerName := strings.ToLower(name)
	if h.hooks[lowerName] != nil {
		for _, handler := range h.hooks[lowerName] {
			handler(args)
		}
	}
}

// Delete will delete all handlers of the hook name supplied
func (h *hookSystem) Delete(name string) {
	lowerName := strings.ToLower(name)
	if h.hooks[lowerName] != nil {
		delete(h.hooks, lowerName)
	}
}
