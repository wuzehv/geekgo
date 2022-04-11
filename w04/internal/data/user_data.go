package data

import (
	"github.com/wuzehv/geekgo/w04/internal/dao"
)

type UserPO struct {
	Id   int
	Name string
	Age  int
}

func (u *UserPO) GetUserList() []UserPO {
	user := dao.User{Name: u.Name, Age: u.Age}
	// 调用dao层代码
	poData := user.SearchUser()

	// dao层数据转PO
	res := make([]UserPO, len(poData))
	for k, v := range poData {
		// 这里没有copy password字段，dao层与PO的隔离
		res[k].Id = v.Id
		res[k].Name = v.Name
		res[k].Age = v.Age
	}

	return res
}
