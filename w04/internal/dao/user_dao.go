package dao

func (u *User) SearchUser() []User {
	// 模拟从数据库获取数据
	c := []User{
		{Id: 1, Name: "user1", Age: 10},
		{Id: 2, Name: "user2", Age: 20},
		{Id: 3, Name: "test", Age: 30},
	}

	res := make([]User, 0)
	for _, v := range c {
		if u.Name == v.Name || u.Age == v.Age {
			res = append(res, v)
		}
	}

	return res
}
