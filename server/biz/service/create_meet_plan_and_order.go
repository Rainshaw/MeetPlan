package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateMeetPlanAndOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateMeetPlanAndOrderService(ctx context.Context, RequestContext *app.RequestContext) *CreateMeetPlanAndOrderService {
	return &CreateMeetPlanAndOrderService{RequestContext: RequestContext, Context: ctx}
}

func (h *CreateMeetPlanAndOrderService) Run(req *model.CreateMeetPlanAndOrderRequest, resp *model.CreateMeetPlanAndOrderResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}