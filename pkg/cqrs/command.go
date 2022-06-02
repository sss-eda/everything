package cqrs

type Command[Payload any] func(Payload) error
