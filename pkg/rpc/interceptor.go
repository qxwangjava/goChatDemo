package rpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return errors.Wrap(err, "")
}

func WrapError(err error) Result {
	r := ERROR
	if err != nil {
		r = Result{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Code:    CODE_FAIL,
		}
		return r
	}
	return r
}

func WrapRPCError(err error) Result {
	r := ERROR
	if err != nil {
		r = Result{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Code:    CODE_FAIL,
		}
		return r
	}
	return r
}
