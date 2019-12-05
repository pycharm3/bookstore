package dao

import (
	_"fmt"
	"bookstore/model"
	"bookstore/utils"
	"strconv"
)

// Getbooks获取books中所有的图书
func Getbooks()([]*model.Book,error){
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	// rows获取到所有的图书信息
	rows,err := utils.Db.Query(sqlStr)
	defer rows.Close()
	if err != nil{
		return nil,err
	}
	var books []*model.Book
	// Next遍历图书信息，循环插入追加信息
	for rows.Next(){
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_Path)
		books = append(books,book)
	}
	return books,nil
}

// 像数据库中添加一本图书
func AddBook(b *model.Book) error{
	sqlstr := "insert into books(title ,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	// 执行sql
	_,err := utils.Db.Exec(sqlstr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.Img_Path)
	if err != nil{
		return err
	}
	return nil
}

// 跟据图书Id删除图书
func DeleteBook(bookId string)error{
	sqlstr := "delete from books where id = ?"
	_,err := utils.Db.Exec(sqlstr,bookId)
	if err != nil{
		return err
	}
	return nil
}

// GetBookId 跟据图书Id从数据库查询出一本书获取要修改的图书
func GetBookId(bookId string)(*model.Book,error){
	// 写sql
	sqlstr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlstr,bookId)
	book := &model.Book{}
	// 为book里的字段赋值
	row.Scan(&book.Id,&book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_Path)
	return book,nil
}

// 跟据图书信息更新图书
func Updatebook(b *model.Book)error{
	// 写sql语句
	sqlstr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	// 执行
	_,err := utils.Db.Exec(sqlstr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.Id)
	if err != nil{
		return err
	}
	return nil
}

// 返回所有图书信息，及分页信息
func GetPageBooks(pageNo string)(*model.Page,error){
	iSales,_ := strconv.ParseInt(pageNo,10,0)

	// 获取数据库中图书总数
	sqlstr := "select count(*) from books"
	// 执行
	row := utils.Db.QueryRow(sqlstr)
	var totalRecord int64
	row.Scan(&totalRecord)
	// 设置每页只显示4条
	var pagesize int64
	pagesize = 4
	var totalPageNo int64
	// 如果总页数取余4刚好为0则页数为总数除以4，否则在其基础加一
	if totalRecord % pagesize == 0{
		totalPageNo = totalRecord / pagesize
	}else{
		totalPageNo = totalRecord / pagesize + 1
	}
	// 获取当前页的图书
	sqlstr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	// 执行
	rows,_ := utils.Db.Query(sqlstr2,(iSales-1)*pagesize,pagesize)
	var books []*model.Book
	for rows.Next(){
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_Path)
		books = append(books,book)
	}
	
	// 创建page
	page := &model.Page{
		Books : books,
		PageNo : iSales,
		PageSize : pagesize,
		TotalPageNo : totalPageNo,
		TotalRecord : totalRecord,
	}
	return page,nil
}

// 返回所有图书信息，及分页信息 带价格范围信息
func GetPageBooksByPrice(pageNo string,minPrice string,maxPrice string)(*model.Page,error){
	iSales,_ := strconv.ParseInt(pageNo,10,0)

	// 获取数据库中图书总数
	sqlstr := "select count(*) from books where price between ? and ?"
	// 执行
	row := utils.Db.QueryRow(sqlstr,minPrice,maxPrice)
	var totalRecord int64
	row.Scan(&totalRecord)
	// 设置每页只显示4条
	var pagesize int64
	pagesize = 4
	var totalPageNo int64
	// 如果总页数取余4刚好为0则页数为总数除以4，否则在其基础加一
	if totalRecord % pagesize == 0{
		totalPageNo = totalRecord / pagesize
	}else{
		totalPageNo = totalRecord / pagesize + 1
	}
	// 获取当前页的图书
	sqlstr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	// 执行
	rows,_ := utils.Db.Query(sqlstr2,minPrice,maxPrice,(iSales-1)*pagesize,pagesize)
	var books []*model.Book
	for rows.Next(){
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_Path)
		books = append(books,book)
	}
	
	// 创建page
	page := &model.Page{
		Books : books,
		PageNo : iSales,
		PageSize : pagesize,
		TotalPageNo : totalPageNo,
		TotalRecord : totalRecord,
	}
	return page,nil
}