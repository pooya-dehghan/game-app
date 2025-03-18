package main

import (
	"fmt"
	"net/http"

	"github.com/pooya-dehghan/entity"
	"github.com/pooya-dehghan/repository/mysql"
)

func main() {
	http.HandleFunc("/users/register", userRegisterHandler)
	http.ListenAndServe(":8787", nil)
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Print(writer, "invalid method")
	}
}

func userRegisterTest() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.RegisterUser(entity.User{
		Name:        "poouya",
		PhoneNumber: "0912992",
	})

	if err != nil {
		fmt.Errorf("some thing went wrong %w", err)
	}
	fmt.Println(createdUser)
}
