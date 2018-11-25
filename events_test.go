package events

import (
	"fmt"
	"os"
	"testing"
)

var events IEventEmitter

func TestMain(m *testing.M) {
	events = CreateEventEmitter()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestBasic(t *testing.T) {
	actual := 0
	expected := 1

	events.On("test", func(args IEventArgs) {
		actual++
	})

	events.Call("test", CreateEventArgs(nil))
	if actual != expected {
		msg := fmt.Sprintf("Expected: %d, Actual: %d", expected, actual)
		t.Error(msg)
	}
}

func TestMutipleHandlers(t *testing.T) {
	actual := 0
	expected := 2

	events.On("test", func(args IEventArgs) {
		actual++
	})

	events.On("test", func(args IEventArgs) {
		actual++
	})

	events.Call("test", CreateEventArgs(nil))
	if actual != expected {
		msg := fmt.Sprintf("Expected: %d, Actual: %d", expected, actual)
		t.Error(msg)
	}
}

func TestDelete(t *testing.T) {
	actual := 0
	expected := 0

	events.On("test", func(args IEventArgs) {
		actual++
	})

	events.Delete("test")
	events.Call("test", CreateEventArgs(nil))
	if actual != expected {
		msg := fmt.Sprintf("Expected: %d, Actual: %d", expected, actual)
		t.Error(msg)
	}
}
