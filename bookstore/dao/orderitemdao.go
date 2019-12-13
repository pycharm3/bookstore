package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 向数据库插入订单项
func AddOrderItem(orderItem *model.OrderItem)error{
	// 写sql
	sqlstr := "insert into order_items(count,amount,title,author,price,img_path,order_id)values(?,?,?,?,?,?,?)"
	_,err := utils.Db.Exec(sqlstr,orderItem.Count,orderItem.Amount,orderItem.Title,orderItem.Author,orderItem.Price,orderItem.ImgPath,orderItem.OrderId)
	if err != nil{
		return err
	}
	return nil
}

// 跟据订单号获取该订单下所有订单项
func GetOrderItemByOrderId(orderId string)([]*model.OrderItem,error){
	sqlstr := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id = ?"
	rows,err := utils.Db.Query(sqlstr,orderId)
	if err != nil{
		return nil,err
	}
	var orderItems []*model.OrderItem
	for rows.Next(){
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemId,&orderItem.Count,&orderItem.Amount,&orderItem.Title,&orderItem.Author,&orderItem.Price,&orderItem.ImgPath,&orderItem.OrderId)
		orderItems = append(orderItems,orderItem)
	}
	return orderItems,nil
}