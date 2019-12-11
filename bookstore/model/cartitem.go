package model 

// 购物项struct
type CartItem struct{
	CartItemId int64 // 购物项id
	Book *Book	// 购物项图书信息book结构体类型
	Count int64	// 购物项中图书数量
	Amount	float64 // 购物项中图书金额小计
	CartId string // 购物车Id
}

// 购物项图书金额小计方法
func (c *CartItem)GetAmount()float64{
	price := c.Book.Price
	// 一本图书数量*单价得到金额
	return float64(c.Count)*price
}
