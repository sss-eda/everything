package everything

type MessageKind string

const (
	CommandMessageKind MessageKind = "command"
	QueryMessageKind   MessageKind = "query"
	EventMessageKind   MessageKind = "event"
)

type Message struct {
	Kind    MessageKind `json:"kind"`
	Payload []byte      `json:"payload"`
}
