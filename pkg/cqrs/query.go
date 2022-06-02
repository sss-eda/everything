package cqrs

type Query[Request, Response any] func(Request) (Response, error)
