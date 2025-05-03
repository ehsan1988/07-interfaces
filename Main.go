package main

import (
	"interface/app"
	"interface/storage"
	"interface/user"
)

func main() {
	application := app.App{

		Name: "my application",
		// StorageFilePath: "user.txt",
		UserStorage: &storage.Memory{},
	}

	myUser := user.User{
		Id:   1,
		Name: "Ehsan Ahmadi",
	}

	application.CreateUser(myUser)

}
