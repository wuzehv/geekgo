package biz

import "github.com/wuzehv/geekgo/w04/internal/data"

type UserDo struct {
	Id int
	Name string
	Age  int
}

func (u *UserDo) GetUserList() []UserDo {
	// 这里面是各种业务逻辑
	// ...

	d := data.UserPO{Name: u.Name, Age: u.Age}
	s := d.GetUserList()

	// PO转DO
	res := make([]UserDo, len(s))
	for k, v := range s {
		res[k].Id = v.Id
		res[k].Name = v.Name
		res[k].Age = v.Age
	}

	// 调用data层代码
	return res
}
