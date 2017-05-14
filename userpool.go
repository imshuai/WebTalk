package main

//UserPool 定义用户池结构
type UserPool struct {
	Users map[string]*User
	count int
}

var users = &UserPool{
	Users: make(map[string]*User, 0),
	count: 0,
}

//Add 向用户池中添加用户
func (up *UserPool) Add(u *User) {
	up.Users[u.NickName] = u
	up.count++
}

//Delete 从用户池中删除用户
func (up *UserPool) Delete(u *User) {
	delete(up.Users, u.NickName)
	up.count--
}

//Count 获得用户池中当前用户数
func (up *UserPool) Count() int {
	return up.count
}
