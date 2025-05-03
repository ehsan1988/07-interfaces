package app

import (
	// "encoding/json"
	// "fmt"
	"interface/user"
	// "os"
)

type UserStore interface {
	CreateUser(u user.User)
	ListUsers() []user.User
	GetUserById(id int) user.User
}

type App struct {
	Name            string
	// StorageFilePath string
	UserStorage UserStore
}

func (app App) CreateUser(u user.User) {
	// f, err := os.OpenFile(app.StorageFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println("cannot open storage file", err)

	// 	return
	// }
	// defer f.Close()

	// //TIP serialize user
	// data, mErr := json.Marshal(u)
	// if mErr != nil {
	// 	fmt.Println("cannot json user", mErr)

	// 	return
	// }
	// _, wErr := f.Write(data)
	// if wErr != nil {
	// 	fmt.Println("cannot write user", wErr)
	// }
	app.UserStorage.CreateUser(u)
}

//func (app App) ListUsers() []User {
//
//}
//func (app App) GetUserById(id int) User {}
