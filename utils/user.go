package utils

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Login    int    `json:"login"`
}

func CreateUser(id string, userName string, login int) User {
	u := User{id, userName, login}
	return u
}

func (u User) getId() string {
	return u.Id
}

func (u User) getUserName() string {
	return u.UserName
}

func (u User) getLogin() int {
	return u.Login
}

func (u *User) setUserName(userName string) {
	u.UserName = userName
}

func (u *User) setLogin(login int) {
	u.Login = login
}
