package controller

import (
	"context"
)

type Request interface {
	Context() context.Context
	Param(name string) string
	Method() string
}

type Response interface {
	SetHeader(name, value string)
	Write(statusCode int, body interface{}) error
}

type Controller interface {
	HandleWindSpeedLoad(req Request, res Response) error
}
