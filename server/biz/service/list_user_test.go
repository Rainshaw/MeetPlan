package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	model "meetplan/biz/model"
	"testing"
)

func TestListUserService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewListUserService(ctx, c)
	// init req and assert value
	req := &model.ListUserRequest{}
	resp := &model.ListUserResponse{}
	err := s.Run(req, resp)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}