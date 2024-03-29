package main

import (
	"net/http"
	"bookstore/controller"
)

func main(){
	/* StripPrefix返回一个处理器，该处理器会将请求的URL.
	Path字段中给定前缀prefix去除后再交由h处理 */
	http.Handle("/pages/",http.StripPrefix("/pages/",http.FileServer(http.Dir("views/pages"))))
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))
	// HandleFunc传入的是一个已实现的处理器handle
	http.HandleFunc("/main",controller.IndexHandler)
	// 登录处理器
	// 必需的 action 属性规定当提交表单时，向何处发送表单数据。
	http.HandleFunc("/login",controller.LoginHandler)
	// 注销处理器
	http.HandleFunc("/logout",controller.LogoutHandler)
	// 注册处理器
	http.HandleFunc("/regist",controller.RegistHandler)
	// 获取所有图书处理器
	// http.HandleFunc("/getBooks",controller.GetBooksHandler)
	// 获取带分页的图书信息
	http.HandleFunc("/getPageBooks",controller.GetPageBooksHandler)
	// 获取图书价格范围
	http.HandleFunc("/getPageBooksByPrice",controller.GetPageBooksByPrice)
	// 添加一本图书
	http.HandleFunc("/addBook",controller.AddBook)
	// 添加图书到购物车
	http.HandleFunc("/addBookTwoCart",controller.AddBookTwoCart)
	// 从购物车删除一本图书
	http.HandleFunc("/deleteCartBook",controller.DeleteCartBook)
	// 清空购物车
	http.HandleFunc("/emptyCart",controller.EmptyCart)
	// 获取购物车信息
	http.HandleFunc("/getCartInfo",controller.GetCartInfo)
	// 删除一本图书
	http.HandleFunc("/deleteBook",controller.DeleteBook)
	// 获取要修改图书资料
	http.HandleFunc("/toUpdateBookPage",controller.ToUpdateBook)
	// 修改图书
	http.HandleFunc("/updateBook",controller.UpdateBook)
	// 结账
	http.HandleFunc("/checkOut",controller.CheckOut)
	// 获取所有订单
	http.HandleFunc("/getOrders",controller.GetAllOrder)
	// 查看订单详情
	http.HandleFunc("/getOrderInfo",controller.GetOrderItemsByOrderId)
	// 查看我的订单
	http.HandleFunc("/getMyOrder",controller.GetMyOrders)
	// 发货
	http.HandleFunc("/sendOrder",controller.SendOrder)
	// 确认收货
	http.HandleFunc("/orderOk",controller.OrderOk)
	// 删除已完成订单
	http.HandleFunc("/delOrder",controller.DelOrder)
	http.ListenAndServe(":8080",nil)
}
