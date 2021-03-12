package rpc

import (
	"context"
	"google.golang.org/grpc"
)

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) Result {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return WrapRPCError(err)
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
