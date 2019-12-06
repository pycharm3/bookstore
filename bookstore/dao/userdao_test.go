package dao

import (
	"testing"
	"fmt"
	"bookstore/model"
)

// func TestUser(t *testing.T){
// 	fmt.Println("测试UserDao中函数")
// 	t.Run("插入user信息",saveuser)
// 	t.Run("验证用户名",chackUsername)
// 	t.Run("验证用户名密码",chackusernameandpassword)
// }

// func TestBook(t *testing.T){
// 	t.Run("测试获取所有图书",getbooks)
// 	t.Run("测试添加一本图书",addBook)
// 	t.Run("测试删除一本图书",deleteBook)
// 	t.Run("测试获取一本图书",getBookId)
// 	t.Run("更新一本图书",updatebook)
// 	t.Run("获取当前页图书信息",getPageBooks)
// }

// func TestSession(t *testing.T){
	// t.Run("测试插入session",addsession)
	// t.Run("测试删除session",delsession)
	// t.Run("测试查询session",getSessionById)

// }

func TestCarts(t *testing.T){
	// t.Run("测试插入购物车",addCartItem)
	// t.Run("测试根据图书id查询购物项",getCartItemByBookId)
	// t.Run("测试根据购物车id查询购物项", getCartItemByCartId)
	t.Run("测试根据用户id查询购物项", getCartByUserId)
}

// 插入信息
func saveuser(t *testing.T){
	err := SaveUser("ming","123456","123@qq.com")
	if err != nil{
		fmt.Println("Saveuser err=",err)
		return
	}else{
		fmt.Println("插入成功")
	}
}

func chackusernameandpassword(t *testing.T){
	user,_ := ChackUsernameAndPassword("tom","123456")
	fmt.Println("根据username和password获取用户信息为: ",*user)
}

func chackUsername(t *testing.T){
	user,err := ChackUsername("tom")
	if err != nil{
		fmt.Println("Saveuser err=",err)
		return
	}
	fmt.Println("根据username查询到的用户为:",*user)
}

// 测试查询图书
func getbooks(t *testing.T){
	books,_ := Getbooks()
	for k,v := range books{
		fmt.Printf("第%v本书为:%v",k+1,*v)
	}
}

// 测试添加一本图书
func addBook(t *testing.T){
	book := &model.Book{
		Title : "三国演义",
		Author : "罗贯中",
		Price : 88.8,
		Sales : 100,
		Stock : 200,
		Img_Path : "/static/img/default.jpg",
	}
	err := AddBook(book)
	if err != nil{
		fmt.Println(err)
		return
	}else{
		fmt.Println("添加成功")
	}
}

func deleteBook(t *testing.T){
	DeleteBook("5")
}

// 测试获取一本图书
func getBookId(t *testing.T){
	book,err := GetBookId("6")
	if err != nil{
		return
	}else{
		fmt.Println(book)
	}
}

// 测试更新图书
func updatebook(t *testing.T){
	book := &model.Book{
		Id : 9,
		Title : "红灯记",
		Author : "卡耐基",
		Price : 88.8,
		Sales : 100,
		Stock : 200,
		Img_Path : "/static/img/default.jpg",
	}
	Updatebook(book)
}

func getPageBooks(t *testing.T){
	page,_ := GetPageBooks("1")
	fmt.Println("当前页数",page.PageNo)
	fmt.Println("总页数",page.TotalPageNo)
	fmt.Println("总记录数",page.TotalRecord)
	fmt.Println("当前页图书有: ",)
	for _,v := range page.Books{
		fmt.Println("图书信息为: ",*v)
	}
}

// 向数据库插入session
func addsession(t *testing.T){
	session := &model.Session{
		SessionId : "132156413113213",
		UserName : "tom",
		UserId : 1,
	}
	err := AddSession(session)
	if err != nil{
		fmt.Println(err)
		return
	}else{
		fmt.Println("session插入成功")
	}
}

// 跟据sessionId查询session
func getSessionById(t *testing.T){
	session,err := GetSessionById("d60fb41b-b246-4e56-42fa-8d6ce8cd5637")
	if err != nil{
		return
	}else{
		fmt.Println("查询成功session = ",*session)
	}
}

// 删除一条session
func delsession(t *testing.T){
	err := DelSession("132156413113213")
	if err != nil{
		return
	}else{
		fmt.Println("session删除成功")
	}
}

// 测试插入购物车 插入购物项
func addCartItem(t *testing.T){
	book := &model.Book{
		Id : 9,
		Price : 10,
	}

	book2 := &model.Book{
		Id : 10,
		Price : 10,
	}

	// 创建一个购物项slice
	var cartItems []*model.CartItem
	// 创建两个购物项
	cartItem := &model.CartItem{
		Book : book,
		Count : 10,
		CartId : 1111,
	}
	cartItems = append(cartItems,cartItem)
	cartItem2 := &model.CartItem{
		Book : book2,
		Count : 10,
		CartId : 1111,
	}
	cartItems = append(cartItems,cartItem2)
	// 创建购物车
	cart := &model.Cart{
		CartId : "1111",
		CartItems : cartItems,
		UserId : 1,
	}
	err := AddCart(cart)
	if err != nil{
		fmt.Println("购物车插入失败",err)
	}else{
		fmt.Println("购物车插入成功")
	}
}

// 测试根据图书id查询当前购物项
func getCartItemByBookId(t *testing.T){
	cartItem,err := GetCartItemByBookId("9")
	if err != nil{
		fmt.Println("查询失败",err)
	}else{
		fmt.Println("查询成功，购物项为: ",*cartItem)
	}
}

// 测试根据购物车id查询当前购物车所有购物项
func getCartItemByCartId(t *testing.T){
	cartItem,err := GetCartItemByCartId("1111")
	if err != nil{
		fmt.Println("查询失败",err)
	}else{
		for _,v := range cartItem{
			fmt.Println("查询成功，购物项为: ",*v)
		}
	}
}

// 测试跟据用户Id查询对应购物车
func getCartByUserId(t *testing.T){
	cart,err := GetCartByUserId(1)
	if err != nil{
		fmt.Println("查询失败，错误信息: ",err)
	}else{
		fmt.Println("查询成功，购物车为: ",cart)
	}
}