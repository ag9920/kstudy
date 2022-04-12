package main

import (
	"context"

	"github.com/ag9920/kstudy/kitex_gen/api"
)

// KStudyImpl implements the last service interface defined in the IDL.
type KStudyImpl struct{}

// Concat implements the KStudyImpl interface.
func (s *KStudyImpl) Concat(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = api.NewResponse()
	resp.Message = req.GetMessage1() + req.GetMessage2()
	return
}
