package controller

import (
	"net/http"
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"text/template"
	_"strconv"
	"fmt"
)

// 添加图书到购物车handle
func AddBookTwoCart(w http.ResponseWriter,q *http.Request){
	// 1，通过session判断用户是否登录
	flag,session := dao.IsLogin(q)
	if flag{
		// 2，获取要添加的图书Id
		bookId := q.FormValue("bookId")
		// 3，跟据图书Id获取图书信息
		book,_ := dao.GetBookId(bookId)
			// 4，通过session信息获取用户id
		userId := session.UserId
		// 5，通过用户Id判断数据库是否有当前用户购物车
		cart,_ := dao.GetCartByUserId(userId)
		// 6，已经有购物车
		if cart != nil{
			// 6.1 通过要插入图书Id及购物车Id获取对应购物项
			cartItem,err := dao.GetCartItemByBookIdAndCartId(bookId,cart.CartId)
			if err != nil{
				fmt.Println("dao.GetCartItemByBookIdAndCartId err = ",err)
			}
			fmt.Println("看看返回结果",cartItem,bookId,cart.CartId)
			if cartItem != nil{
				// 6.2 如果有该图书对应购物项将购物项加一
				// 6.2.1 获取购物车切片里的所有购物项
				cst := cart.CartItems
				for _,v := range cst{
					// 6.2.2 获取购物车/购物项切片/购物项
					if v.Book.Id == cartItem.Book.Id{
						// 6.2.3 将购物项中count加一
						v.Count += 1
						// 6.2.4 通过cartItem更新数据库中对应购物项的数量及总金额
						err := dao.UpdateBookCount(v)
						if err != nil{
							fmt.Println("dao.UpdateBookCount err = ",err)
						}
					}
				}
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
			err = dao.UpdateCart(cart)
			if err != nil{
				fmt.Println("dao.UpdateCart err = ",err)
			}
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
		// 11 发送给前端页面,通过异步请求将书名添加时显示
		w.Write([]byte("您刚刚将 ["+book.Title+"] 添加到购物车!"))
	}else{
		w.Write([]byte("请先登录"))
	}
}

// 跟据用户Id获取购物车信息 
func GetCartInfo(w http.ResponseWriter,q *http.Request){
	_,session := dao.IsLogin(q)
	// 获取用户id
	userId := session.UserId
	// 跟据用户Id从数据库获取对应购物车
	cart,_ := dao.GetCartByUserId(userId)
	if cart != nil{
		// 解析模板
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		// 将cart赋给session.Cart
		session.Cart = cart
		// 将session返回给前端页面
		t.Execute(w,session)
	}else{
		// 该用户还没有购物车
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w,session)
	}
}

// 从购物车中删除一本图书
func DeleteCartBook(w http.ResponseWriter,q *http.Request){
	cartItemId := q.FormValue("cartItemId")
	cartItem,err := dao.GetCartItemById(cartItemId)
	if err != nil{
		fmt.Println("dao.GetCartItemById err",err)
	}
	if cartItem.Count > 1{
		cartItem.Count -= 1
		dao.UpdateBookCount(cartItem)
	}else{
		dao.DelCartItemById(cartItemId)
	}
	flag,session := dao.IsLogin(q)
	if flag{
		userId := session.UserId
		cart,err := dao.GetCartByUserId(userId)
		if err != nil{
			fmt.Println("GetCartByUserId(userId) err = ",err)
		}
		dao.UpdateCart(cart)
	}
	GetCartInfo(w,q)
}

// 清空购物车
func EmptyCart(w http.ResponseWriter,q *http.Request){
	cartId := q.FormValue("cartId")
	err := dao.EmptyCart(cartId)
	if err != nil{
		fmt.Println("dao.EmptyCart err = ",err)
	}

	flag,session := dao.IsLogin(q)
	if flag{
		userId := session.UserId
		cart,err := dao.GetCartByUserId(userId)
		if err != nil{
			fmt.Println("GetCartByUserId(userId) err = ",err)
		}
		dao.UpdateCart(cart)
	}
	GetCartInfo(w,q)
}