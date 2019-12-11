package model

type Session struct{
	SessionId string
	UserName string
	UserId int
	Cart *Cart	// 在session中增添一个Cart字段用于购物车判断
}