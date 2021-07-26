package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const connString = "jscx:Jxdb#2020@tcp(172.29.49.241:3306)/hxchain?charset=utf8&parseTime=True&loc=Asia%2FShanghai"

func main() {
	/** 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
	//以上作业，要求提交到自己的 GitHub 上面，然后把自己的 GitHub 地址填写到班班提供的表单中：
	//https://jinshuju.net/f/D3jwL8
	*/
	// invoke query
	name, pwd, err1 := queryDB("1")
	name2, pwd2, err2 := queryDB("1234")
	fmt.Printf("uid=1 name:%s,pwd:%s,error:%+v\n", name, pwd, err1)
	fmt.Printf("uid=1234 name:%s,pwd:%s,error:%+v\n", name2, pwd2, err2)

}

func queryDB(uid string) (name string, pwd string, err error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Printf("连接失败：%s", err)
	}
	// invoke query
	row := db.QueryRow("select * from user where uid=?", uid)
	errScan := row.Scan(&uid, &name, &pwd)
	fmt.Println("查找id=10的数据：", uid, ", ", name, "------", pwd)
	if errScan != nil {
		fmt.Printf("main:34 -> 查询出现异常：%s\n", errScan)
		return "", "", errors.New("查无记录")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("close连接失败：%s\n", err)
		}
		fmt.Println("关闭数据库")
	}(db)
	return name, pwd, nil
}
