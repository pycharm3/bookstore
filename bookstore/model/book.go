package model

type Book struct{
	Id int	// 图书Id
	Title string	// 图书名
	Author string	// 图书作者
	Price float64	// 图书金额
	Sales int	// 图书销量
	Stock int	// 图书库存
	Img_Path string	// 图书照片路径
}