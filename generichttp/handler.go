package main

import (
	"context"

	"github.com/ag9920/kstudy/generichttp/kitex_gen/http"
)

// BizServiceImpl implements the last service interface defined in the IDL.
type BizServiceImpl struct{}

// BizMethod1 implements the BizServiceImpl interface.
func (s *BizServiceImpl) BizMethod1(ctx context.Context, req *http.BizRequest) (resp *http.BizResponse, err error) {
	// TODO: Your code here...
	resp = &http.BizResponse{
		ItemCount: []int64{324, 32},
	}
	return
}

// BizMethod2 implements the BizServiceImpl interface.
func (s *BizServiceImpl) BizMethod2(ctx context.Context, req *http.BizRequest) (resp *http.BizResponse, err error) {
	// TODO: Your code here...
	return
}

// BizMethod3 implements the BizServiceImpl interface.
func (s *BizServiceImpl) BizMethod3(ctx context.Context, req *http.BizRequest) (resp *http.BizResponse, err error) {
	// TODO: Your code here...
	return
}
