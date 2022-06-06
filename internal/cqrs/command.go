package cqrs

type Command func(Payload) error
