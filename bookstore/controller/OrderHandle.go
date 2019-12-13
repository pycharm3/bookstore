package controller

import (
	"fmt"
	"net/http"
	"bookstore/dao"
	"text/template"
	"bookstore/utils"
	"bookstore/model"
	"time"
)

// 去结账
func CheckOut(w http.ResponseWriter,q *http.Request){
	// 1 通过请求信息获取session
	_,session := dao.IsLogin(q)

	// 2 通过session获取用户Id
	userId := session.UserId

	// 3 通过用户Id获取当前用户购物车
	cart,_ := dao.GetCartByUserId(userId)

	// 4 将购物车信息设置到订单中
	orderId := utils.CreateUUID()

	// 4.1 创建订单生成时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	
	order := &model.Order{
		OrderId : orderId,
		CreateTime : timeStr,
		TotalCount : cart.GetTotalCount(),
		TotalAmount : cart.GetTotalAmount(),
		State : 0,
		UserId : int64(cart.UserId),
	}

	// 5 调用dao函数将订单信息入库
	err := dao.AddOrder(order)
	if err != nil{
		fmt.Println("dao.AddOrder err = ",err)
	}

	// 6 遍历得到cart中的cartItem
	for _,v := range cart.CartItems{
		// 6.1 将cartItem信息存入OrderItem实例
		orderItem := &model.OrderItem{
			Count : v.Count,
			Amount : v.GetAmount(),
			Title : v.Book.Title,
			Author : v.Book.Author,
			Price : v.Book.Price,
			ImgPath : v.Book.Img_Path,
			OrderId : orderId,
		}

		// 6.2 将orderItem入库
		err := dao.AddOrderItem(orderItem)
		if err != nil{
			fmt.Println("dao.AddOrderItem err = ",err)
		}

		// 6.3 订单完成后更新当前图书库存和销量
		book := v.Book
		book.Sales += int(v.Count)	// 图书销量加一
		book.Stock -= int(v.Count) // 图书库存减一

		// 6.4 更新图书信息
		err = dao.Updatebook(book)
		if err != nil{
			fmt.Println("dao.Updatebook err",err)
		}
	}

	// 7 订单,订单项入库完毕清空购物车
	dao.EmptyCart(cart.CartId)
	// 8 将order设置到session中用于前端显示
	session.Order = order

	// 9 解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w,session)
}

// 获取所有订单
func GetAllOrder(w http.ResponseWriter,q *http.Request){
	orders,_ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w,orders)
}

// 跟据订单号(orderId)获取该订单下的所有订单项
func GetOrderItemsByOrderId(w http.ResponseWriter,q *http.Request){
	orderId := q.FormValue("OrderId")
	orderItems,err := dao.GetOrderItemByOrderId(orderId)
	if err != nil{
		fmt.Println("dao.GetOrderItemByOrderId(orderId) err = ",err)
	}
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w,orderItems)
}

// 跟据当前用户Id获取该用户下的所有订单
func GetMyOrders(w http.ResponseWriter,q *http.Request){
	_,session := dao.IsLogin(q)
	// 2 通过session获取用户Id
	userId := session.UserId
	orders,err := dao.GetMyOrders(userId)
	// 将订单设置到session
	session.Orders = orders
	if err != nil{
		fmt.Println("dao.GetMyOrders err = ",err)
	}
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w,session)
}

// 跟据orderId修改对应state值实现发货
func SendOrder(w http.ResponseWriter,q *http.Request){
	orderId := q.FormValue("orderId")
	err := dao.UpdateOrderStateOne(orderId)
	if err != nil{
		fmt.Println("dao.UpdateOrderStateOne err = ",err)
	}
	GetAllOrder(w,q)
}

// 跟据orderId修改对应state值实现确认收货
func OrderOk(w http.ResponseWriter,q *http.Request){
	orderId := q.FormValue("orderId")
	err := dao.UpdateOrderStateTwo(orderId)
	if err != nil{
		fmt.Println("dao.UpdateOrderState err = ",err)
	}
	GetMyOrders(w,q)
}

// 跟据orderId删除order
func DelOrder(w http.ResponseWriter,q *http.Request){
	orderId := q.FormValue("orderId")
	err := dao.DelOrder(orderId)
	if err != nil{
		fmt.Println("dao.DelOrder err",err)
	}
	GetMyOrders(w,q)
}