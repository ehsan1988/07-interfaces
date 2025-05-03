package storage

import "interface/user"

type Memory struct {
	Users []user.User
}

func (m *Memory) AddUser(u user.User) {
	m.Users = append(m.Users, u)
}

func (m *Memory) CreateUser(u user.User) {
	m.Users = append(m.Users, u)
}
func (m *Memory) ListUsers() []user.User       {}
func (m *Memory) GetUserById(id int) user.User {}
