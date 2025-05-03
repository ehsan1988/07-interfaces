package app

import (
	"encoding/json"
	"fmt"
	"interface/user"
	"os"
)

type App struct {
	Name            string
	StorageFilePath string
}

func (app App) CreateUser(u user.User) {
	f, err := os.OpenFile(app.StorageFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("cannot open storage file", err)

		return
	}
	defer f.Close()

	//TIP serialize user
	data, mErr := json.Marshal(u)
	if mErr != nil {
		fmt.Println("cannot json user", mErr)

		return
	}
	_, wErr := f.Write(data)
	if wErr != nil {
		fmt.Println("cannot write user", wErr)
	}
}

//func (app App) ListUsers() []User {
//
//}
//func (app App) GetUserById(id int) User {}
