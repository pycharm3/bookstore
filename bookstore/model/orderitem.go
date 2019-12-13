package model


type OrderItem struct{
	OrderItemId  int64	// 订单项Id，自增
	Count		 int64	// 订单项中图书的数量
	Amount		 float64	// 订单项中图书金额小计
	Title		 string		// 订单项中图书的书名
	Author		 string		// 订单项中图书的作者
	Price		 float64	// 订单项中图书的价格
	ImgPath		 string		// 订单项中图书的封面
	OrderId		 string		// 订单项所属的订单
}