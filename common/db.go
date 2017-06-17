package common

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)
func Query(stmt string)(datas []map[string]string,err error){
	db,err := getDB()
	defer db.Close()
	if err!= nil{
		fmt.Println("获取数据库连接失败",err)
	}
	rows,err := db.Query(stmt)
	if err != nil {
		fmt.Println("Query Error...", err)
		return nil,err
	}
	defer rows.Close()
	cols,err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	values := make([]sql.RawBytes,len(cols))
	scanArgs := make([]interface{},len(values))
	for i := range values{
		scanArgs[i]=&values[i]
	}

	for rows.Next(){
		err = rows.Scan(scanArgs...)
		var value string
		data := map[string]string{}
		for i,col := range values{
			if col == nil{
				value ="null"
			}else{
				value = string(col)
			}
			colu:=string(cols[i])
			data[colu]=value
		}
		datas = append(datas,data)
	}

	return datas,nil
}

func getDB()(db *sql.DB,err error){
	db,err = sql.Open("mysql","root:123456@tcp(localhost:3306)/full_marketing")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	return db,nil
}

