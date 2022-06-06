package cqrs

type Query func(Request) (Response, error)

type Request interface{}

type Response interface{}
