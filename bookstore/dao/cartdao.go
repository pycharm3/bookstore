package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 向数据库中插入购物车
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

// GetCartByUserId 跟据用户Id查找对应购物车
func GetCartByUserId(uesrId int)(*model.Cart,error){
	sqlstr := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	row := utils.Db.QueryRow(sqlstr,uesrId)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartId,&cart.TotalCount,&cart.TotalAmount,&cart.UserId)
	if err != nil{
		return nil,err
	}
	// 获取当前购物车中所有购物项
	cartItems,_ := GetCartItemByCartId(cart.CartId)
	// 将所有购物项设置到购物车中
	cart.CartItems = cartItems
	return cart,nil
}

// 更新购物车的总数量和总金额
func UpdateCart(cart *model.Cart)error{
	// 写sql
	sqlstr := "update carts set total_count = ? , total_amount = ? where id = ?"
	// 执行
	_,err := utils.Db.Exec(sqlstr,cart.GetTotalCount(),cart.GetTotalAmount(),cart.CartId)
	if err != nil{
		return err
	}
	return nil
}