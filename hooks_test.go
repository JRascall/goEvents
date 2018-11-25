package hooks

import (
	"fmt"
	"os"
	"testing"
)

var events IHookSys

func TestMain(m *testing.M) {
	events = CreateHookSystem()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestBasic(t *testing.T) {
	actual := 0
	expected := 1

	events.On("test", func(args IHookArgs) {
		actual++
	})

	events.Call("test", nil)
	if actual != expected {
		msg := fmt.Sprintf("Expected: %d, Actual: %d", expected, actual)
		t.Error(msg)
	}
}

func TestMutipleHandlers(t *testing.T) {
	actual := 0
	expected := 2

	events.On("test", func(args IHookArgs) {
		actual++
	})

	events.On("test", func(args IHookArgs) {
		actual++
	})

	events.Call("test", nil)
	if actual != expected {
		msg := fmt.Sprintf("Expected: %d, Actual: %d", expected, actual)
		t.Error(msg)
	}
}

func TestDelete(t *testing.T) {
	actual := 0
	expected := 0

	events.On("test", func(args IHookArgs) {
		actual++
	})

	events.Delete("test")
	events.Call("test", nil)
	if actual != expected {
		msg := fmt.Sprintf("Expected: %d, Actual: %d", expected, actual)
		t.Error(msg)
	}
}

func TestArgs(t *testing.T) {
	var actual interface{}
	_ = actual
	expected := 1

	events.On("test", func(args IHookArgs) {
		data := args.Data()

		actual := data["testVal"]
		if actual != expected {
			msg := fmt.Sprintf("Expected: %d, Actual: %d", expected, actual)
			t.Error(msg)
		}
	})

	args := make(map[string]interface{})
	args["testVal"] = 1
	events.Call("test", CreateHookArgs(args))
}
