package model

// 购物车struct
type Cart struct{
	CartId	string			// 购物id
	CartItems	[]*CartItem	// 购物车中所有的购物项
	TotalCount	int64		// 购物车中图书总数，通过计算得到
	TotalAmount	 float64	// 购物车图书总金额，计算得到
	UserId	int	 			// 购物车所属用户id
}

// GetTotalCount 获取购物车图书总数量
func (c *Cart)GetTotalCount()int64{
	var totalCount int64
	for _,v := range c.CartItems{
		totalCount = totalCount + v.Count
	}
	return totalCount
}

// GetTotalAmount 获取购物车图书金额总和
func (c *Cart)GetTotalAmount()int64{
	var totalAmount int64
	for _,v := range c.CartItems{
		totalAmount = totalAmount + v.Amount
	}
	return totalAmount
}
