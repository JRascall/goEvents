package events

type IEventArgs interface {
	Data() map[string]interface{}
}
