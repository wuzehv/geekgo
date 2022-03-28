package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// NameNotFound 包装底层的错误
var NameNotFound = fmt.Errorf("name not found %w", sql.ErrNoRows)

// getName 模拟dao层方法
// 如果没有查到数据，返回wrap的错误变量
func getName(id int) error {
	err := sql.ErrNoRows
	if err == sql.ErrNoRows {
		return NameNotFound
	}

	return err
}

func main() {
	// 模拟业务调用
	err := getName(100)
	// 这里使用我们的错误来处理，好处是对持久化的底层实现没有依赖
	if errors.Is(err, NameNotFound) {
		fmt.Println(err)
		return
	}

	if err != nil {
		log.Fatalln(err)
	}
}