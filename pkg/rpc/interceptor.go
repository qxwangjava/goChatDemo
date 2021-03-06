package rpc

import (
	"context"
	"goChatDemo/pkg/gerror"
	"google.golang.org/grpc"
)

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func WrapError(err error) gerror.Result {
	r := gerror.ERROR
	if err != nil {
		r = gerror.Result{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Code:    gerror.CODE_FAIL,
		}
		return r
	}
	return r
}

func WrapRPCError(err error) gerror.Result {
	r := gerror.ERROR
	if err != nil {
		r = gerror.Result{
			Success: false,
			Message: err.Error(),
			Data:    nil,
			Code:    gerror.CODE_FAIL,
		}
		return r
	}
	return r
}
