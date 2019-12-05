package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 向购物车表中插入购物车
func AddCart(cart *model.Cart)error{
	//将购物车插入到数据库中
	sqlstr := "insert into carts(id,total_count,total_amount,user_id)values(?,?,?,?)"
	_,err := utils.Db.Exec(sqlstr,cart.CartId,cart.GetTotalCount(),cart.GetTotalAmount(),cart.UserId)
	if err != nil{
		return err
	}

	// 获取购物车中所有购物项
	cartItems := cart.CartItems
	// 遍历得到每一个购物项
	for _,v := range cartItems{
		// 将购物车中的购物项插入到数据库中
		AddCartItem(v)
	}
	return nil
}