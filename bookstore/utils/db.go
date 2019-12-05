package utils

import (
	"fmt"
	"database/sql"
	_"mysql-master"
)

var (
	Db *sql.DB
	err error
)

func init(){
	Db,err = sql.Open("mysql","root:password@tcp(localhost:3306)/golang")
	if err != nil{
		fmt.Println("sql.open err = ",err)
		return
	}
}
