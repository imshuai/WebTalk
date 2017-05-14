package main

type UserPool struct {
	Users map[string]*User
	count int
}

func (up *UserPool) Add(u *User) {
	up.Users[u.NickName] = u
	up.count++
}

func (up *UserPool) Delete(u *User) {
	delete(up.Users, u.NickName)
	up.count--
}

func (up *UserPool) Count() int {
	return up.count
}
