package controller

import (
	"net/http"
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	_"text/template"
	_"strconv"
	"fmt"
	
)

// 添加图书到购物车handle
func AddBookTwoCart(w http.ResponseWriter,q *http.Request){
	// 1，获取要添加的图书Id
	bookId := q.FormValue("bookId")
	// 2，跟据图书Id获取图书信息
	book,_ := dao.GetBookId(bookId)
	// 3，通过session判断用户是否登录
	_,session := dao.IsLogin(q)
	// 4，通过session信息获取用户id
	userId := session.UserId
	// 5，通过用户Id判断数据库是否有当前用户购物车
	cart,_ := dao.GetCartByUserId(userId)
	// 6，已经有购物车
	if cart != nil{
		// 6.1 通过要插入图书Id及购物车Id获取对应购物项
		cartItem,_ := dao.GetCartItemByBookIdAndCartId(bookId,cart.CartId)
		fmt.Println("就想看看",cartItem,bookId,cart.CartId)
		if cartItem != nil{
			// 6.2 如果有该图书对应购物项将购物项加一
		}else{
			// 6.3 购物车中还没有该图书对应购物项则创建一个购物项
			cartitem := &model.CartItem{
				Book : book,
				Count : 1,
				CartId : cart.CartId,
			}
			// 6.4 将新创建的购物项追加到该用户的购物车购物项切片中
			cart.CartItems = append(cart.CartItems,cartitem)
			// 6.5 将该购物项添加到当前用户购物车的购物项切片中
			dao.AddCartItem(cartitem)
		}
		// 6.6 不管当前购物车是否有要添加的图书的购物项都要更新数据库购物车中图书的总数量和总金额
		dao.UpdateCart(cart)
	}else{
		// 7 当前用户还没有购物车，创建购物车
		cartId := utils.CreateUUID()
		cart := &model.Cart{
			CartId : cartId,
			UserId : userId, 
		}
		// 8 创建购物车中的购物项切片及购物项实例
		var cartitems []*model.CartItem
		cartitem := &model.CartItem{
			Book : book,
			Count : 1,
			CartId : cartId,
		}
		// 9 将购物项追加到购物车的购物项切片中
		cartitems = append(cartitems,cartitem)
		cart.CartItems = cartitems
		// 10 将购物车插入到数据库
		err := dao.AddCart(cart)
		if err != nil{
			// 插入失败
		}
	}
}