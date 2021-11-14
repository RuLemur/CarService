package transport

import "github.com/go-kit/kit/endpoint"

type EndpointsDescription struct {
	Endpoint            endpoint.Endpoint
	Request             interface{}
	Response            interface{}
}
