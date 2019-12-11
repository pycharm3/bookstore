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

//GetCartItemByBookIDAndCartID 根据图书的id和购物车的id获取对应的购物项
func GetCartItemByBookIdAndCartId(bookId string, cartId string) (*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	//执行	
	row := utils.Db.QueryRow(sqlStr, bookId, cartId)
	//设置一个变量接收图书的id
	//创建cartItem
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
	if err != nil {
		fmt.Println("row.Scan err = ",err)
		return nil, err
	}
	//根据图书的id查询图书信息
	book,_ := GetBookId(bookId)
	//将book设置到购物项
	cartItem.Book = book
	return cartItem, nil
}

// 跟据传入cartItem更新购物项中图书数量和总金额
func UpdateBookCount(cartItem *model.CartItem)error{
	// 写sql语句
	sqlstr := "update cart_items set count = ? ,amount = ? where book_id = ? and cart_id = ?"
	_,err := utils.Db.Exec(sqlstr,cartItem.Count,cartItem.GetAmount(),cartItem.Book.Id,cartItem.CartId)
	if err != nil{
		return err
	}
	return nil
}

// GetCartItemByCartId 根据购物车Id把购物车中所有购物项查询出来
func GetCartItemByCartId(cartId string)([]*model.CartItem,error){
	// 写sql跟据购物车id获取该购物车下所有购物项
	sqlstr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	// 执行sql
	rows,err := utils.Db.Query(sqlstr,cartId)
	if err != nil{
		return nil,err
	}
	var cartItems []*model.CartItem
	for rows.Next(){
		// 新建一个变量接收book_id
		var bookId string
		// 创建cartItem
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemId,&cartItem.Count,&cartItem.Amount,&bookId,&cartItem.CartId)
		if err != nil{
			return nil,err
		}
		// 将获取到的book赋给cartItem.Book
		book,_ := GetBookId(bookId)
		cartItem.Book = book
		cartItems = append(cartItems,cartItem)
	}
	return cartItems,nil
}

// 跟据购物项Id获取购物项
func GetCartItemById(cartItemId string)(*model.CartItem,error){
	sqlstr := "select * from cart_items where id = ?"
	row := utils.Db.QueryRow(sqlstr, cartItemId)
	cartItem := &model.CartItem{}
	var bookId string
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count,&cartItem.Amount,&bookId,&cartItem.CartId)
	if err != nil {
		fmt.Println("row.Scan err = ",err)
		return nil, err
	}
	book,_ := GetBookId(bookId)
	cartItem.Book = book
	return cartItem,nil
}

// 跟据购物项Id删除购物项
func DelCartItemById(cartItemId string)error{
	sqlstr := "delete from cart_items where id = ?"
	_,err := utils.Db.Exec(sqlstr,cartItemId)
	if err != nil{
		return err
	}
	return nil
}

// 清空购物车
func EmptyCart(cartId string)error{
	sqlstr := "delete from cart_items where cart_id = ?"
	_,err := utils.Db.Exec(sqlstr,cartId)
	if err != nil{
		return err
	}
	return nil
}