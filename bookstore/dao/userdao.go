package dao

import (
	"fmt"
	"bookstore/model"
	"bookstore/utils"
)

// 根据用户名和密码从数据库中查询数据
func ChackUsernameAndPassword(username string,password string)(*model.User,error){
	sqlstr := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlstr,username,password)
	user := &model.User{}
	row.Scan(&user.Id,&user.Username,&user.Password,&user.Email)
	return user,nil
}

// 根据用户名从数据库中查询数据
func ChackUsername(username string)(*model.User,error){
	sqlstr := "select id,username,password,email from users where username = ?"
	row := utils.Db.QueryRow(sqlstr,username)
	user := &model.User{}
	row.Scan(&user.Id,&user.Username,&user.Password,&user.Email)
	return user,nil
}

// 向数据库插入信息
func SaveUser(username string,password string,email string)error{
	sqlstr := "insert into users(username,password,email) values(?,?,?)"
	_,err := utils.Db.Exec(sqlstr,username,password,email)
	if err != nil{
		fmt.Println("utils.Db.Exce() err = ",err)
		return err
	}
	return nil
}
