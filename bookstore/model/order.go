package model

// 订单结构体
type Order struct{
	OrderId string	// 订单Id
	CreateTime string	// 订单生成时间	
	TotalCount	int64	// 订单中图书数量
	TotalAmount  float64	// 订单中总金额
	State	int64	// 订单状态 0 未发货，1 已发货，2 已签收
	UserId	int64	// 订单所属用户Id
	UserName string // 用户名
	StateTwo int64 // 订单状态2 表示已发货点击确认收货
	StateThree int64 // 订单状态3 表示订单完成
}