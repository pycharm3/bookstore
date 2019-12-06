// 向数据库中添加session
package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"net/http"
)

// 添加session
func AddSession(session *model.Session)error{
	sqlstr := "insert into sessions values(?,?,?)"
	_,err := utils.Db.Exec(sqlstr,session.SessionId,session.UserName,session.UserId)
	if err != nil{
		return err
	}
	return nil
}

// 查询session
func GetSessionById(sessionId string)(*model.Session,error){
	sqlstr := "select session_id,username,user_id from sessions where session_id = ?"
	// 预编译
	stmt,err := utils.Db.Prepare(sqlstr)
	if err != nil{
		return nil,err
	}
	row := stmt.QueryRow(sessionId)
	// 实例化一个session用来存放查询出的数据
	session := &model.Session{}
	row.Scan(&session.SessionId,&session.UserName,&session.UserId)
	return session,nil
}

// 删除session
func DelSession(sessionId string)error{
	sqlstr := "delete from sessions where session_id = ?"
	_,err := utils.Db.Exec(sqlstr,sessionId)
	if err != nil{
		return err
	}
	return nil
}

// 判断用户是否登录
func IsLogin(q *http.Request)(bool,*model.Session){
	cookie,_ := q.Cookie("user")
	if cookie != nil{
		// 获取cookie的value
		cookieValue := cookie.Value
		// 根据cookieValue查对应的sesson
		session,_ := GetSessionById(cookieValue)
		if session.UserId > 0{
			// 已经登录
			return true,session
		}
	}
	// 没有登录
	return false,nil
}