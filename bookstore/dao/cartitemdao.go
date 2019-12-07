package dao

import (
	"bookstore/utils"
	"bookstore/model"
	"fmt"
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

// GetCartItemByBookId 根据图书Id和购物车Id获取对应购物项
func GetCartItemByBookIdAndCartId(bookId string,cartId string)(*model.CartItem,error){
	// 写sql
	sqlstr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	
	// 执行sql
	fmt.Printf("看看sql,%T,%T,%T",sqlstr,bookId,cartId)
	row := utils.Db.QueryRow(sqlstr, bookId, cartId)
	// 创建CartItem实例接收查询结果
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId,&cartItem.Count,&cartItem.Amount,&cartItem.CartId)
	if err != nil{
		return nil,err
	}
	fmt.Println(cartItem)
	return cartItem,nil
}

// GetCartItemByCartId 根据购物车Id把购物车中所有购物项查询出来
func GetCartItemByCartId(cartId string)([]*model.CartItem,error){
	// 写sql
	sqlstr := "select id,count,amount,cart_id from cart_items where cart_id = ?"
	// 执行sql
	rows,err := utils.Db.Query(sqlstr,cartId)
	if err != nil{
		return nil,err
	}
	var cartItems []*model.CartItem
	for rows.Next(){
		// 创建cartItem
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemId,&cartItem.Count,&cartItem.Amount,&cartItem.CartId)
		if err != nil{
			return nil,err
		}
		cartItems = append(cartItems,cartItem)
	}
	return cartItems,nil
}
