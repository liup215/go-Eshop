package common

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)

func GetDB()(db *sql.DB,err error){
	db,err = sql.Open("mysql","root:123456@tcp(localhost:3306)/full_marketing")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	return db,nil
}
