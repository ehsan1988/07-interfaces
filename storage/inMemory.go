package storage

import "interface/user"

type Memory struct {
	Users []user.User
}

func (m *Memory) AddUser(u user.User) {
	m.Users = append(m.Users, u)
}
