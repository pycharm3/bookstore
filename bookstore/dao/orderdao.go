package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 向数据库中插入订单order
func AddOrder(order *model.Order)error{
	// 写sql
	sqlstr := "insert into orders(id,create_time,total_count,total_amount,state,user_id)values(?,?,?,?,?,?)"
	// 执行
	_,err := utils.Db.Exec(sqlstr,order.OrderId,order.CreateTime,order.TotalCount,order.TotalAmount,order.State,order.UserId)
	if err != nil{
		return err
	}
	return nil
}

// 获取数据库所有Order
func GetOrders()([]*model.Order,error){
	// 写sql
	sqlstr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	// 执行sql
	rows,err := utils.Db.Query(sqlstr)
	if err != nil{
		return nil,err
	}
	var orders []*model.Order
	for rows.Next(){
		order := &model.Order{}
		rows.Scan(&order.OrderId,&order.CreateTime,&order.TotalCount,&order.TotalAmount,&order.State,&order.UserId)
		
		if order.State == 1{
			order.StateTwo = 1
		}
		if order.State == 2{
			order.StateThree = 3
		}

		orders = append(orders,order)
	}
	return orders,nil
}

// 获取当前用户的Order
func GetMyOrders(userId int)([]*model.Order,error){
	sqlstr := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"
	rows,err := utils.Db.Query(sqlstr,userId)
	if err != nil{
		return nil,err
	}
	var orders []*model.Order
	for rows.Next(){
		order := &model.Order{}
		rows.Scan(&order.OrderId,&order.CreateTime,&order.TotalCount,&order.TotalAmount,&order.State,&order.UserId)
		
		if order.State == 1{
			order.StateTwo = 1
		}
		if order.State == 2{
			order.StateThree = 3
		}

		orders = append(orders,order)
	}
	return orders,nil
}


// 跟据orderId修改对应order的state字段实现发货(将state修改为1)
func UpdateOrderStateOne(orderId string)error{
	sqlstr := "update orders set state = 1 where id = ?"
	_,err := utils.Db.Exec(sqlstr,orderId)
	if err != nil{
		return err
	}
	return nil
}

// 跟据orderId修改对应order的state字段实现确认收货
func UpdateOrderStateTwo(orderId string)error{
	sqlstr := "update orders set state = 2 where id = ?"
	_,err := utils.Db.Exec(sqlstr,orderId)
	if err != nil{
		return err
	}
	return nil
}


// 跟据orderId删除order
func DelOrder(orderId string)error{
	sqlstr := "delete from order_items where order_id = ?"
	_,err := utils.Db.Exec(sqlstr,orderId)
	sqlstr = "delete from orders where id = ?"
	_,err = utils.Db.Exec(sqlstr,orderId)
	if err != nil{
		return err
	}
	return nil
}