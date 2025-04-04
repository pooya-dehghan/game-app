package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pooya-dehghan/repository/mysql"
	"github.com/pooya-dehghan/service/authservice"
	"github.com/pooya-dehghan/service/userservice"
)

const SIGNED_KEY = "jwt_key"

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

	userSvc := userservice.New(mysqlRep, []byte(SIGNED_KEY))

	userCreated, err := userSvc.Register(reqData)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in user creation %s`, err.Error())))
		return
	}

	fmt.Println(userCreated)

	userMarshalData, err := json.Marshal(userCreated)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in marshaling response %s`, err.Error())))
		return
	}
	writer.Write(userMarshalData)

}

func loginHandler(writer http.ResponseWriter, req *http.Request) {
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

	authSvc := authservice.NewService("SignKey", "as", "fs", mysqlRep, time.Minute*24, time.Hour*24)

	authRes, err := authSvc.Login(reqData)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in user creation %s`, err.Error())))
		return
	}

	fmt.Println(authRes)

	authReponse, err := json.Marshal(authRes)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in marshaling response %s`, err.Error())))
		return
	}

	writer.Write(authReponse)
}

func ProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Print(writer, "invalid method")
	}

	pReq := userservice.ProfileRequest{UserID: 0}

	mysqlRep := mysql.New()

	uSvc := userservice.New(mysqlRep, []byte(SIGNED_KEY))

	pRes, err := uSvc.Profile(pReq)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`unexpected error %s`, err.Error())))
		return
	}

	data, err := json.Marshal(pRes)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in marshaling %s`, err.Error())))
		return
	}

	writer.Write(data)
}
