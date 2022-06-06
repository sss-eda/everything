package lemi025

type CommandKind string

const (
	ReadConfigCommandKind CommandKind = "read-config"
	ReadTimeCommandKind   CommandKind = "read-time"
	SetTimeCommandKind    CommandKind = "set-time"
)

type Command struct {
	Kind    CommandKind `json:"kind"`
	Payload []byte      `json:"payload"`
}

type ReadConfigCommand struct{}
