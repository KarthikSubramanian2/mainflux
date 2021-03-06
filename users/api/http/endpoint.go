package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mainflux/mainflux/users"
)

func registrationEndpoint(svc users.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(userReq)

		if err := req.validate(); err != nil {
			return nil, err
		}

		err := svc.Register(req.user)
		return tokenRes{}, err
	}
}

func loginEndpoint(svc users.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(userReq)

		if err := req.validate(); err != nil {
			return nil, err
		}

		token, err := svc.Login(req.user)
		if err != nil {
			return nil, err
		}

		return tokenRes{token}, nil
	}
}
