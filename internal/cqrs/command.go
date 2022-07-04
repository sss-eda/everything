package cqrs

// type Command func(Payload) error

// type CommandHandler[Command interface{ Type() CommandType }] func(Command) error

type CommandKind string

type CommandHandler func(command any) error
