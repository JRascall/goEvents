package hooks

type IHookArgs interface {
	Data() map[string]interface{}
}
