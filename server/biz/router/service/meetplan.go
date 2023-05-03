// Code generated by hertz generator. DO NOT EDIT.

package service

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	service "meetplan/biz/handler/service"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			_v1.POST("/friendlink", append(_createfriendlinkMw(), service.CreateFriendLink)...)
			_v1.GET("/friendlink", append(_listfriendlinkMw(), service.ListFriendLink)...)
			_v1.GET("/login", append(_loginMw(), service.Login)...)
			_v1.POST("/meetplan", append(_createmeetplanMw(), service.CreateMeetPlan)...)
			_v1.DELETE("/meetplan", append(_deletemeetplansMw(), service.DeleteMeetPlans)...)
			_v1.POST("/meetplanorder", append(_createmeetplanandorderMw(), service.CreateMeetPlanAndOrder)...)
			_v1.POST("/order", append(_createorderMw(), service.CreateOrder)...)
			_v1.PUT("/termdate", append(_updatetermdaterangeMw(), service.UpdateTermDateRange)...)
			_v1.GET("/termdate", append(_gettermdaterangeMw(), service.GetTermDateRange)...)
			_v1.GET("/updaterecord", append(_listupdaterecordMw(), service.ListUpdateRecord)...)
			_v1.POST("/updaterecord", append(_createupdaterecordMw(), service.CreateUpdateRecord)...)
			_v1.POST("/user", append(_createuserMw(), service.CreateUser)...)
			_v1.GET("/meetplan", append(_listmeetplanMw(), service.ListMeetPlan)...)
			_meetplan := _v1.Group("/meetplan", _meetplanMw()...)
			_meetplan.PUT("/:id", append(_updatemeetplanMw(), service.UpdateMeetPlan)...)
			_meetplan.DELETE("/:id", append(_deletemeetplanMw(), service.DeleteMeetPlan)...)
			{
				_meetplan0 := _v1.Group("/meetplan", _meetplan0Mw()...)
				_meetplan0.GET("/:id", append(_getmeetplanMw(), service.GetMeetPlan)...)
			}
			_v1.GET("/order", append(_listorderMw(), service.ListOrder)...)
			_order := _v1.Group("/order", _orderMw()...)
			_order.PUT("/:id", append(_updateorderMw(), service.UpdateOrder)...)
			{
				_order0 := _v1.Group("/order", _order0Mw()...)
				_order0.GET("/:id", append(_getorderMw(), service.GetOrder)...)
			}
			_v1.GET("/user", append(_listuserMw(), service.ListUser)...)
			_user := _v1.Group("/user", _userMw()...)
			_user.PUT("/:id", append(_updateuserMw(), service.UpdateUser)...)
			{
				_user0 := _v1.Group("/user", _user0Mw()...)
				_user0.GET("/:id", append(_getuserMw(), service.GetUser)...)
				_user0.GET("/self", append(_getselfMw(), service.GetSelf)...)
			}
		}
	}
}
