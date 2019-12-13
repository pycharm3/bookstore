package controller

import (
	"net/http"
	"bookstore/dao"
	"bookstore/model"
	"text/template"
	"strconv"
)

// 写一个去首页的处理器handler
// ResponseWriter接口被HTTP处理器用于构造HTTP回复。
// Request类型代表一个服务端接受到的或者客户端发送出去的HTTP请求。
func IndexHandler(w http.ResponseWriter,q *http.Request){
	// 获取页码
	pageNo := q.FormValue("pageNo")
	if pageNo == ""{
		pageNo = "1"
	}
	// 获取带分页的图书函数
	page,_ := dao.GetPageBooks(pageNo)
	// 调用IsLogin判断用户是否登录成功过
	flag,session := dao.IsLogin(q)
		if flag{
			// 已经登录 设置page里的两个字段IsLogin Username
			page.IsLogin = true
			page.Username = session.UserName
		}

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w,page)
}

// 获取所有图书
func GetBooksHandler(w http.ResponseWriter,q *http.Request){
	Books,_ := dao.Getbooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w,Books)
}

// 带分页获取所有图书
func GetPageBooksHandler(w http.ResponseWriter,q *http.Request){
	// 获取页码
	pageNo := q.FormValue("pageNo")
	if pageNo == ""{
		pageNo = "1"
	}
	// 获取带分页的图书函数
	page,_ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w,page)
}

// 添加一本图书
func AddBook(w http.ResponseWriter,q *http.Request){
	// 跟据name获取图书信息
	title := q.PostFormValue("title")
	price := q.PostFormValue("price")
	author := q.PostFormValue("author")
	sales := q.PostFormValue("sales")
	stock := q.PostFormValue("stock")
	// 将价格销量库存进行转换
	// 返回一个字符串表示的float64类型值
	fprice,_ := strconv.ParseFloat(price,64)
	// 返回字符串表示的整数值，接受正负号。base进制为10进制
	iSales,_ := strconv.ParseInt(sales,10,0)
	iStock,_ := strconv.ParseInt(stock,10,0)

	book := &model.Book{
		Title : title,
		Author : author,
		Price : fprice,
		Sales : int(iSales),
		Stock : int(iStock),
		Img_Path : "/static/img/default.jpg",
	}
	dao.AddBook(book)
	// 添加完查询
	// GetBooksHandler(w,q)
	GetPageBooksHandler(w,q)
}

// 删除一本图书
func DeleteBook(w http.ResponseWriter,q *http.Request){
	bookId := q.FormValue("bookId")
	// 调用删除函数
	err := dao.DeleteBook(bookId)
	if err != nil{
		return
	}
	// GetBooksHandler(w,q)
	GetPageBooksHandler(w,q)
}

// 获取要修改图书
func ToUpdateBook(w http.ResponseWriter,q *http.Request){
	bookId := q.FormValue("bookId")
	book,_ := dao.GetBookId(bookId)
	t := template.Must(template.ParseFiles("views/pages/manager/book_modify.html"))
	t.Execute(w,book)
}

// 修改图书
func UpdateBook(w http.ResponseWriter,q *http.Request){
	// 跟据name获取图书信息
	bookId := q.PostFormValue("bookId")
	title := q.PostFormValue("title")
	price := q.PostFormValue("price")
	author := q.PostFormValue("author")
	sales := q.PostFormValue("sales")
	stock := q.PostFormValue("stock")
	// 将价格销量库存进行转换
	// 返回一个字符串表示的float64类型值
	fprice,_ := strconv.ParseFloat(price,64)
	// 返回字符串表示的整数值，接受正负号。base进制为10进制
	iSales,_ := strconv.ParseInt(sales,10,0)
	iStock,_ := strconv.ParseInt(stock,10,0)
	isid,_ := strconv.ParseInt(bookId,10,0)

	book := &model.Book{
		Id : int(isid),
		Title : title,
		Author : author,
		Price : fprice,
		Sales : int(iSales),
		Stock : int(iStock),
		Img_Path : "/static/img/default.jpg",
	}
	dao.AddBook(book)
	// 添加完查询
	GetPageBooksHandler(w,q)
}


// 获取筛选价格信息
func GetPageBooksByPrice(w http.ResponseWriter,q *http.Request){
	var page *model.Page
	pageNo := q.FormValue("pageNo")
	if pageNo == ""{
		pageNo = "1"
	}
	min := q.FormValue("min")
	max := q.FormValue("max")

	if min == "" && max == ""{
		page,_ = dao.GetPageBooks(pageNo)
	}else{
		page,_ = dao.GetPageBooksByPrice(pageNo,min,max)
		// 将价格范围设置到page中
		page.MinPrice = min
		page.MaxPrice = max
	}

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w,page)
}

