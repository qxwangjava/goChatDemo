package rpc

import (
	"context"
	"goChatDemo/pkg/web"

	"google.golang.org/grpc"
)

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) web.Result {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return WrapRPCError(err)
}

func WrapError(err error) web.Result {
	r := web.ERROR
	if err != nil {
		r = web.Result{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Code:    web.CODE_FAIL,
		}
		return r
	}
	return r
}

func WrapRPCError(err error) web.Result {
	r := web.ERROR
	if err != nil {
		r = web.Result{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Code:    web.CODE_FAIL,
		}
		return r
	}
	return r
}
