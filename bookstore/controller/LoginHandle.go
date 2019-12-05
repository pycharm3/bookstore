package controller

import (
	"fmt"
	"net/http"
	"bookstore/dao"
	"text/template"
	"bookstore/utils"
	"bookstore/model"
)

// 处理用户登录请求
func LoginHandler(w http.ResponseWriter,q *http.Request){
	// 登录成功后如果再次发送请求直接跳转到首页
	flag,_ := dao.IsLogin(q)
	if flag{
		IndexHandler(w,q)
	}else{
		// 获取用户名和密码
		username := q.PostFormValue("username")
		password := q.PostFormValue("password")
		// 调用userdao进行验证
		user,_ := dao.ChackUsernameAndPassword(username,password)
		if user.Id > 0{
			// 生成uuid用做sessionId
			uuid := utils.CreateUUID()
			// 创建一个session实例
			session := &model.Session{
				SessionId : uuid,
				UserName : user.Username,
				UserId : user.Id,
			}
			// 将登录成功的用户session存入数据库中
			dao.AddSession(session)
			// 创建一个cookie与session关联
			cookie := http.Cookie{
				Name : "user",
				Value : uuid,
				HttpOnly : true,
			}
			// 将cookie发送给浏览器
			http.SetCookie(w,&cookie)

			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w,user)
		}else{
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w,"用户名或密码不正确")
		}
	}	
}

// 处理用户注销请求
func LogoutHandler(w http.ResponseWriter,q *http.Request){
	// 获取用户cookie
	cookie,_ := q.Cookie("user")
	if cookie != nil{
		cookieValue := cookie.Value
		// 删除该cookie数据库对应的session
		dao.DelSession(cookieValue)
		// 设置cookie失效 MaxAge<0表示立刻删除该cookie，等价于"Max-Age: 0"
		cookie.MaxAge = -1
		// 将修改后的cookie发送给浏览器 cookie已经是指针类型
		http.SetCookie(w,cookie)
	}
	// 注销后跳转到首页
	IndexHandler(w,q)
}

// 处理用户注册请求 向数据库插入请求
func RegistHandler(w http.ResponseWriter,q *http.Request){
	// 获取用户名
	username := q.PostFormValue("username")
	password := q.PostFormValue("password")
	email := q.PostFormValue("email")
	// 调用userdao进行验证
	user,_ := dao.ChackUsername(username)
	if user.Id == 0{
		err := dao.SaveUser(username,password,email)
		if err != nil{
			fmt.Println("dao.SaveUser err",err)
			return
		}
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w,"")
	}else{
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w,"用户名已存在")
	}
}