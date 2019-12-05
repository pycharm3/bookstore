package dao

import (
	"bookstore/utils"
	"bookstore/model"
)

// 向数据库中插入购物项cartItem
func AddCartItem(cartItem *model.CartItem)error{
	sqlstr := "insert into cart_items(count,amount,book_id,cart_id)values(?,?,?,?)"
	// 执行sql
	_,err := utils.Db.Exec(sqlstr,cartItem.Count,cartItem.GetAmount(),cartItem.Book.Id,cartItem.CartId)
	if err != nil{
		return err
	}
	return nil
}