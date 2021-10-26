package helpers

import (
	"context"
	"errors"
	"time"
)

type handler interface {
	Create(ctx context.Context, body []byte) (Response, error)
}

const fiveSecondsTimeout = time.Second * 15

func Router(handler handler) func(context.Context, Request) (Response, error) {
	return func(ctx context.Context, req Request) (Response, error) {

		ctx, cancel := context.WithTimeout(ctx, fiveSecondsTimeout)
		defer cancel()

		switch req.HTTPMethod {

		case "POST":
			return handler.Create(ctx, []byte(req.Body))

		default:
			return Response{}, errors.New("invalid method")
		}
	}
}
