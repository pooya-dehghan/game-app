package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pooya-dehghan/repository/mysql"
	"github.com/pooya-dehghan/service/authservice"
	"github.com/pooya-dehghan/service/userservice"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", loginHandler)
	http.ListenAndServe(":8787", mux)
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Print(writer, "invalid method")
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in read body %s`, err.Error())))
		return
	}
	var reqData userservice.RegisterRequest
	err = json.Unmarshal(data, &reqData)
	fmt.Println(reqData)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in read body %s`, err.Error())))
		return
	}
	mysqlRep := mysql.New()

	userSvc := userservice.New(mysqlRep)

	userCreated, err := userSvc.Register(reqData)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in user creation %s`, err.Error())))
		return
	}

	fmt.Println(userCreated)
}

func loginHandler(writer http.ResponseWriter, req *http.Request) {
	//passes the password and phone number
	if req.Method != http.MethodPost {
		fmt.Print(writer, "invalid method")
	}

	data, err := io.ReadAll(req.Body)

	var reqData authservice.LoginRequest
	err = json.Unmarshal(data, &reqData)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in read body %s`, err.Error())))
		return
	}

	mysqlRep := mysql.New()

	authSvc := authservice.NewService(mysqlRep)

	authRes, err := authSvc.Login(reqData)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in user creation %s`, err.Error())))
		return
	}

	fmt.Println(authRes)
}
