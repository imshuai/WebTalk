package main

import (
	"errors"
	"time"
)

//User 定义用户信息
type User struct {
	ID            int64
	NickName      string    `xorm:"unique"`
	IsOnline      bool      `xorm:"default(0)"`
	CreateTime    time.Time `xorm:"created"`
	LastLoginTime time.Time
	Friends       []int64
	Session       map[string]string `xorm:"-"`
}

//NewUser 新建一个用户
func NewUser(nickname string) *User {
	u := &User{
		NickName:      nickname,
		IsOnline:      false,
		LastLoginTime: time.Now(),
	}
	_, err := db.InsertOne(u)
	if err != nil {
		LogError("insert new user into database fail with error:", err)
		return nil
	}
	LogDebug("insert new user into database success with id:", u.ID)
	return u
}

//UserOnline 处理用户上线
func UserOnline(nickname string) *User {
	u := new(User)
	exist, err := db.Where("`nick_name`=?", nickname).Get(u)
	if err != nil {
		LogError("get user from database fail with error:", err)
		return nil
	}
	if !exist {
		LogInfo("get user", nickname, "from database fail, user does not exist")
		return nil
	}
	u.IsOnline = true
	u.LastLoginTime = time.Now()
	_, err = db.AllCols().Update(u)
	if err != nil {
		LogError("update user into database fail with error:", err)
		return nil
	}
	return u
}

//SendMessageTo 发送消息到指定的用户
func (u *User) SendMessageTo(nickname string, msg []byte) (bool, error) {
	t := new(User)
	exist, err := db.Where("`nick_name`=?", nickname).Get(t)
	if err != nil {
		LogError("get user from database fail with error:", err)
		return false, err
	}
	if !exist {
		LogInfo("get user", nickname, "from database fail, user does not exist")
		return false, errors.New("send message to " + nickname + " fail, user does not exist")
	}
	if !t.reciveMsg(msg) {
		return false, errors.New("user " + nickname + "cannot recive message, maybe he is not online")
	}
	return true, nil
}

func (u *User) reciveMsg(msg []byte) bool {
	return false
}
