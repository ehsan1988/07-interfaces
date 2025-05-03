package main

import (
	"interface/app"
	"interface/user"
)

func main() {
	application := app.App{

		Name:            "my application",
		StorageFilePath: "user.txt",
	}

	myUser := user.User{
		Id:   1,
		Name: "Ehsan Ahmadi",
	}

	application.CreateUser(myUser)

}
