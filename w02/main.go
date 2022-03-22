package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// getName 模拟dao层方法
// 如果没有查到数据，返回wrap的sql.ErrNoRows，提供原生sql
func getName(id int, name *string) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	querySql := "select name from t1 where id = ?"
	err = db.QueryRow(querySql, id).Scan(name)
	if err == sql.ErrNoRows {
		return fmt.Errorf("%w, sql: %s, bind: %d", err, querySql, id)
	}

	return err
}

func main() {
	var name string
	// 模拟业务调用
	err := getName(100, &name)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err)
		return
	}

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(name)
}