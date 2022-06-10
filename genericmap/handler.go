package main

import (
	"context"
	"fmt"

	server0 "github.com/ag9920/kstudy/genericmap/kitex_gen/kitex/test/server"
)

// ExampleServiceImpl implements the last service interface defined in the IDL.
type ExampleServiceImpl struct{}

// ExampleMethod implements the ExampleServiceImpl interface.
func (s *ExampleServiceImpl) ExampleMethod(ctx context.Context, req *server0.ExampleReq) (resp *server0.ExampleResp, err error) {
	fmt.Println("receive req", req)
	// TODO: Your code here...
	resp = &server0.ExampleResp{
		Msg: "this is from Example Method",
	}
	return
}
