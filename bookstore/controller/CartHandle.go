package controller

import (
	"net/http"
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	_"text/template"
	_"strconv"
	
)

// 添加图书到购物车handle
func AddBookTwoCart(w http.ResponseWriter,q *http.Request){
	// 获取要添加的图书Id
	bookId := q.FormValue("bookId")
	// 跟据图书Id获取图书信息
	book,_ := dao.GetBookId(bookId)
	// 判断用户是否登录
	_,session := dao.IsLogin(q)
	// 获取用户id
	userId := session.UserId
	// 判断数据库中是否有当前用户的购物车
	cart,_ := dao.GetCartByUserId(userId)
	if cart != nil{
		// 当前用户已经有购物车
	}else{
		// 当前用户还没有购物车
		// 1 创建购物车
		// 生成购物车的id
		cartId := utils.CreateUUID()
		cart := &model.Cart{
			CartId : cartId,
			UserId : userId, // 这里我有点疑惑，userId应该是空的
		}
		// 创建购物车中的购物项
		// 声明一个购物项切片
		var cartitems []*model.CartItem
		cartitem := &model.CartItem{
			Book : book,
			Count : 1,
			CartId : cartId,
		}
		cartitems = append(cartitems,cartitem)
		cart.CartItems = cartitems
		// 将购物车插入到数据库
		err := dao.AddCart(cart)
		if err != nil{
			// 插入失败
		}
	}
}